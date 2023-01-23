package gateway

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Goscord/goscord/goscord/gateway/packet"
	"github.com/gorilla/websocket"
	"net"
	"strconv"
	"sync"
	"time"
)

type VoiceConnection struct {
	sync.RWMutex

	session *Session
	ready   bool

	UserId    string
	GuildId   string
	ChannelId string

	deaf         bool
	mute         bool
	speaking     bool
	reconnecting bool

	// ws connection
	connMu sync.Mutex
	conn   *websocket.Conn

	// udp connection
	udpConn *net.UDPConn

	// Voice connection data
	sessionId string
	token     string
	endpoint  string
	ip        string
	port      int
	ssrc      uint32
	modes     []string

	// recv/send channels (maybe use an io.Reader/Writer instead?)

	close chan struct{}
}

func (v *VoiceConnection) login() error {
	v.Lock()
	defer v.Unlock()

	i := 0
	for {
		if v.sessionId != "" {
			break
		}

		if i > 20 {
			return fmt.Errorf("cannot resolve session id")
		}

		time.Sleep(50 * time.Millisecond)

		i++
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://"+v.endpoint+"/?v=4", nil)
	if err != nil {
		return errors.New("cannot connect to packet server: " + err.Error())
	}

	v.conn = conn

	voiceIdentify := packet.NewVoiceIdentify(v.GuildId, v.UserId, v.sessionId, v.token)
	if err := v.Send(voiceIdentify); err != nil {
		return errors.New("cannot send identify packet: " + err.Error())
	}

	v.close = make(chan struct{})

	go v.listen(v.conn, v.close)

	return nil
}

func (v *VoiceConnection) listen(conn *websocket.Conn, close <-chan struct{}) {
	for {
		_, message, err := v.conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, 4014) {
				v.Lock()
				v.conn = nil
				v.Unlock()

				for i := 0; i < 5; i++ {
					<-time.After(1 * time.Second)

					v.RLock()
					reconnected := v.conn != nil
					v.RUnlock()
					if !reconnected {
						continue
					}
					return
				}

				v.session.Lock()
				delete(v.session.VoiceConnections, v.GuildId)
				v.session.Unlock()

				v.Close()

				return
			}

			v.RLock()
			sameConnection := v.conn == conn
			v.RUnlock()

			if sameConnection {
				go v.reconnect()
			}

			return
		}

		select {
		case <-close:
			return
		default:
			go v.handleEvent(message)
		}
	}
}

func (v *VoiceConnection) handleEvent(msg []byte) {
	pk, err := packet.NewPacket(msg)

	if err != nil {
		return
	}

	switch pk.Opcode {
	case packet.OpVoiceReady:
		var ready packet.VoiceReady
		if err := json.Unmarshal(msg, &ready); err != nil {
			return
		}

		v.Lock()
		v.ready = true
		v.ip = ready.Data.IP
		v.port = ready.Data.Port
		v.ssrc = ready.Data.SSRC
		v.modes = ready.Data.Modes
		v.Unlock()

		err := v.loginUDP()
		if err != nil {
			fmt.Println(err)
			return
		}

	case packet.OpVoiceHello:
		var hello packet.VoiceHello
		if err := json.Unmarshal(msg, &hello); err != nil {
			return
		}

		interval := time.Duration(hello.Data.HeartbeatInterval) // little hack cuz discord sends a float64

		go v.startHeartbeat(v.conn, v.close, interval)

	case packet.OpVoiceSessionDescription:
		// TODO: Retrieve encryption key

		// TODO: Handle other voice event
	}
}

func (v *VoiceConnection) loginUDP() error {
	v.Lock()
	defer v.Unlock()

	if v.conn == nil {
		return errors.New("nil connection")
	}

	if v.udpConn != nil {
		return errors.New("udp connection already exists")
	}

	if v.close == nil {
		return errors.New("nil close channel")
	}

	if v.endpoint == "" {
		return errors.New("empty endpoint")
	}

	host := v.ip + ":" + strconv.Itoa(v.port)
	addr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return err
	}

	v.udpConn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	buf := make([]byte, 70)
	binary.BigEndian.PutUint32(buf, v.ssrc)
	_, err = v.udpConn.Write(buf)
	if err != nil {
		return err
	}

	buf = make([]byte, 70)
	bufLen, _, err := v.udpConn.ReadFromUDP(buf)
	if err != nil {
		return err
	}

	if bufLen < 70 {
		return errors.New("invalid udp response")
	}

	// read ip and port from ip discovery packet
	ip := string(buf[4:20])
	port := binary.BigEndian.Uint16(buf[68:70])

	voiceSelect := packet.NewVoiceSelectProtocol(ip, port)
	if err := v.Send(voiceSelect); err != nil {
		return errors.New("cannot send select protocol packet: " + err.Error())
	}

	// TODO: Try to heartbeat UDP

	return nil
}

func (v *VoiceConnection) startHeartbeat(wsConn *websocket.Conn, close <-chan struct{}, i time.Duration) {
	if close == nil || wsConn == nil {
		return
	}

	var _ error

	ticker := time.NewTicker(i * time.Millisecond)
	defer ticker.Stop()

	for {
		if err := v.Send(packet.NewVoiceHeartbeat(time.Now().UnixNano())); err != nil {
			// TODO: Log error
			return
		}

		select {
		case <-ticker.C:
			// nothing

		case <-close:
			return
		}
	}
}

func (v *VoiceConnection) wait() error {
	i := 0

	for {
		v.RLock()
		ready := v.ready
		v.RUnlock()

		if ready {
			return nil
		}

		if i > 10 {
			return fmt.Errorf("cannot resolve voice connection")
		}

		time.Sleep(1 * time.Second)

		i++
	}
}

func (v *VoiceConnection) reconnect() {
	fmt.Println("reconnecting") // TODO: Remove debug

	v.Lock()
	if v.reconnecting {
		v.Unlock()

		return
	}

	v.reconnecting = true
	v.Unlock()

	defer func() {
		v.Lock()
		v.reconnecting = false
		v.Unlock()
	}()

	v.Close()

	wait := time.Duration(1)

	for {
		<-time.After(wait * time.Second)
		wait *= 2
		if wait > 600 {
			wait = 600
		}

		if v.session.Status() < StatusReady || v.session.conn == nil {
			continue
		}

		_, err := v.session.JoinVoiceChannel(v.GuildId, v.ChannelId, v.mute, v.deaf)
		if err == nil {
			return
		}

		payload := packet.NewVoiceStateUpdate(v.GuildId, "", true, true)

		v.session.Send(payload)
	}
}

func (v *VoiceConnection) Close() {
	v.Lock()
	defer v.Unlock()

	v.ready = false
	v.speaking = false

	if v.close != nil {
		close(v.close)
		v.close = nil
	}

	if v.udpConn != nil {
		if err := v.udpConn.Close(); err != nil {
			// TODO: Log error
		}

		v.udpConn = nil
	}

	if v.conn != nil {
		v.connMu.Lock()
		if err := v.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
			// TODO: Log error
		}
		v.connMu.Unlock()

		time.Sleep(1 * time.Second)

		if err := v.conn.Close(); err != nil {
			// TODO: Log error
		}

		v.conn = nil
	}
}

func (s *VoiceConnection) Send(v interface{}) error {
	s.connMu.Lock()
	defer s.connMu.Unlock()

	return s.conn.WriteJSON(v)
}

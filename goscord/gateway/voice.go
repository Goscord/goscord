package gateway

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Goscord/goscord/goscord/gateway/packet"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"net"
	"strconv"
	"sync"
	"time"
)

type VoiceConnection struct {
	sync.RWMutex

	session *Session
	ready   atomic.Bool

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
	v.RLock()
	sessionId := v.sessionId
	endpoint := v.endpoint
	guildId := v.GuildId
	userId := v.UserId
	token := v.token
	v.RUnlock()

	v.connMu.Lock()
	defer v.connMu.Unlock()

	attempt := 0
	for {
		v.RLock()
		sessionId := v.sessionId
		v.RUnlock()

		if sessionId != "" {
			break
		}

		if attempt > 20 {
			return errors.New("failed to login to voice")
		}

		<-time.After(50 * time.Millisecond)

		attempt++
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://"+endpoint, nil)
	if err != nil {
		return errors.New("cannot connect to voice websocket server")
	}

	v.conn = conn

	payload := packet.NewVoiceIdentify(guildId, userId, sessionId, token)
	if err := conn.WriteJSON(payload); err != nil {
		return err
	}

	v.Lock()
	cclose := make(chan struct{})
	v.close = cclose
	v.Unlock()

	go v.listen(conn, cclose)

	return nil
}

func (v *VoiceConnection) listen(conn *websocket.Conn, close chan struct{}) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, 4014) {
				v.Lock()
				v.conn = nil
				v.Unlock()

				for i := 0; i < 5; i++ {
					<-time.After(1 * time.Second)

					v.connMu.Lock()
					reconnected := v.conn != nil
					v.connMu.Unlock()
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

			v.connMu.Lock()
			sameConnection := v.conn == conn
			v.connMu.Unlock()

			if sameConnection {
				go v.reconnect()
			}

			return
		}

		select {
		case <-close:
			return

		default:
			pk, err := packet.NewPacket(msg)

			if err != nil {
				return
			}

			fmt.Println("[VOICE] Opcode: ", pk.Opcode)

			switch pk.Opcode {
			case packet.OpVoiceReady:
				var ready packet.VoiceReady
				if err := json.Unmarshal(msg, &ready); err != nil {
					return
				}

				v.ready.Store(true)

				v.Lock()
				v.ip = ready.Data.IP
				v.port = ready.Data.Port
				v.ssrc = ready.Data.SSRC
				v.modes = ready.Data.Modes
				v.Unlock()

				err := v.loginUdp()
				if err != nil {
					return
				}

			case packet.OpVoiceHello:
				var hello packet.VoiceHello
				if err := json.Unmarshal(msg, &hello); err != nil {
					return
				}

				interval := time.Duration(hello.Data.HeartbeatInterval) // little hack cuz discord sends a float64

				v.RLock()
				c := v.close
				v.RUnlock()

				v.connMu.Lock()
				conn := v.conn
				v.connMu.Unlock()

				go v.startHeartbeat(conn, c, interval)

			case packet.OpVoiceSessionDescription:
				// TODO: Retrieve encryption key

				// TODO: Handle other voice event
			}

		}
	}
}

func (v *VoiceConnection) loginUdp() error { // TODO: make this work
	v.connMu.Lock()
	conn := v.conn
	v.connMu.Unlock()

	v.RLock()
	udpConn := v.udpConn
	c := v.close
	endpoint := v.endpoint
	vIp := v.ip
	vPort := v.port
	v.RUnlock()

	if conn == nil {
		return errors.New("nil connection")
	}

	if udpConn != nil {
		return errors.New("udp connection already exists")
	}

	if c == nil {
		return errors.New("nil close channel")
	}

	if endpoint == "" {
		return errors.New("empty endpoint")
	}

	host := vIp + ":" + strconv.Itoa(vPort)
	addr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return err
	}

	udpConn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	v.Lock()
	v.udpConn = udpConn
	v.Unlock()

	buf := make([]byte, 70)
	binary.BigEndian.PutUint32(buf, v.ssrc)
	_, err = udpConn.Write(buf)
	if err != nil {
		return err
	}

	buf = make([]byte, 70)
	bufLen, _, err := udpConn.ReadFromUDP(buf)
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

	go v.heartbeatUdp(udpConn, c)

	return nil
}

func (v *VoiceConnection) heartbeatUdp(udpConn *net.UDPConn, close <-chan struct{}) { // TODO: make this work
	if udpConn == nil || close == nil {
		return
	}

	payload := make([]byte, 8)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {

		binary.LittleEndian.PutUint32(payload, 0x80)
		binary.LittleEndian.PutUint32(payload, 0xC8)

		_, err := udpConn.Write(payload)
		if err != nil {
			return
		}

		select {
		case <-ticker.C:
			// loop

		case <-close:
			return
		}
	}
}

func (v *VoiceConnection) startHeartbeat(conn *websocket.Conn, c <-chan struct{}, i time.Duration) {
	if c == nil || conn == nil {
		return
	}

	ticker := time.NewTicker(i * time.Millisecond)
	defer ticker.Stop()

	for {
		if err := v.Send(packet.NewVoiceHeartbeat(time.Now().UnixMilli())); err != nil {
			// TODO: Log error
			return
		}

		select {
		case <-ticker.C:
			// nothing

		case <-c:
			return
		}
	}
}

func (v *VoiceConnection) wait() error {
	var attempt int = 1
	var err error

	for {
		if v.ready.Load() {
			break
		}

		if attempt > 10 {
			err = errors.New("voice connection timed out")
			break
		}

		<-time.After(1 * time.Second)

		attempt++
	}

	return err
}

func (v *VoiceConnection) reconnect() {
	v.RLock()
	if v.reconnecting {
		v.RUnlock()

		return
	}

	v.Lock()
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

		v.RLock()
		session := v.session
		guildId := v.GuildId
		channelId := v.ChannelId
		mute := v.mute
		deaf := v.deaf
		v.RUnlock()

		if session.Status() != StatusReady {
			// TODO: Log error
			continue
		}

		_, err := session.JoinVoiceChannel(guildId, channelId, mute, deaf)
		if err == nil {
			// TODO: Log success
			return
		}

		payload := packet.NewVoiceStateUpdate(guildId, "", false, false)
		session.Send(payload)
	}
}

func (v *VoiceConnection) Speaking(speaking bool) error { return nil }

func (v *VoiceConnection) Disconnect() (err error) {
	v.RLock()
	sessionId := v.sessionId
	guildId := v.GuildId
	v.RUnlock()

	if sessionId != "" {
		payload := packet.NewVoiceStateUpdate(guildId, "", false, false)
		err = v.session.Send(payload)

		v.Lock()
		v.sessionId = ""
		v.Unlock()
	}

	v.Close()

	v.session.Lock()
	delete(v.session.VoiceConnections, guildId)
	v.session.Unlock()

	return
}

func (v *VoiceConnection) Close() {
	v.ready.Store(false)

	v.Lock()
	v.speaking = false
	v.Unlock()

	v.RLock()
	c := v.close
	udpConn := v.udpConn
	v.RUnlock()

	if c != nil {
		v.Lock()
		close(v.close)
		v.close = nil
		v.Unlock()
	}

	if udpConn != nil {
		err := udpConn.Close()
		if err != nil {
			// TODO: Log error
		}

		v.Lock()
		v.udpConn = nil
		v.Unlock()
	}

	v.connMu.Lock()
	defer v.connMu.Unlock()

	if v.conn != nil {
		err := v.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			// TODO: Log error
		}

		<-time.After(1 * time.Second)

		err = v.conn.Close()
		if err != nil {
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

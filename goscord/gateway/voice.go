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

func (v *VoiceConnection) login() error { return nil }

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
		go v.startHeartbeat(v.conn, v.close, interval)
		v.RUnlock()

	case packet.OpVoiceSessionDescription:
		// TODO: Retrieve encryption key

		// TODO: Handle other voice event
	}
}

func (v *VoiceConnection) loginUdp() error {
	v.RLock()
	defer v.RUnlock()

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

	go v.heartbeatUdp(v.udpConn, v.close)

	return nil
}

func (v *VoiceConnection) heartbeatUdp(udpConn *net.UDPConn, close <-chan struct{}) {
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

func (v *VoiceConnection) startHeartbeat(wsConn *websocket.Conn, close <-chan struct{}, i time.Duration) {
	if close == nil || wsConn == nil {
		return
	}

	var _ error

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
			return fmt.Errorf("voice connection timed out")
		}

		<-time.After(1 * time.Second)

		i++
	}
}

func (v *VoiceConnection) reconnect() {
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

		if v.session.Status() != StatusReady {
			// TODO: Log error
			continue
		}

		_, err := v.session.JoinVoiceChannel(v.GuildId, v.ChannelId, v.mute, v.deaf)
		if err == nil {
			// TODO: Log success
			return
		}

		payload := packet.NewVoiceStateUpdate(v.GuildId, "", false, false)
		v.session.Send(payload)
	}
}

func (v *VoiceConnection) Speaking(speaking bool) error { return nil }

func (v *VoiceConnection) Disconnect() error { return nil }

func (v *VoiceConnection) Close() {}

func (s *VoiceConnection) Send(v interface{}) error {
	s.connMu.Lock()
	defer s.connMu.Unlock()

	return s.conn.WriteJSON(v)
}

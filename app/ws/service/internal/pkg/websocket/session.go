package websocket

import (
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
)

var channelBufSize = 256

type SessionID string

type Session struct {
	id     SessionID
	conn   *ws.Conn
	send   chan []byte
	server *Server
}

func (s *Session) SessionId() SessionID {
	return s.id
}

func NewSession(conn *ws.Conn, server *Server) *Session {
	if conn == nil {
		panic("websocket connect cannot be nil")
	}
	uid, _ := uuid.NewUUID()

	return &Session{
		id:     SessionID(uid.String()),
		conn:   conn,
		send:   make(chan []byte, channelBufSize),
		server: server,
	}
}

func (s *Session) Listen() {
	go s.writePump()
	go s.readPump()
}

func (s *Session) SendTextMessage(message []byte) error {
	if s.conn == nil {
		return nil
	}
	return s.conn.WriteMessage(ws.TextMessage, message)
}

func (s *Session) sendPingMessage(message string) error {
	if s.conn == nil {
		return nil
	}
	return s.conn.WriteMessage(ws.PingMessage, []byte(message))
}

func (s *Session) sendPongMessage(message string) error {
	if s.conn == nil {
		return nil
	}
	return s.conn.WriteMessage(ws.PongMessage, []byte(message))
}

func (s *Session) sendBinaryMessage(message []byte) error {
	if s.conn == nil {
		return nil
	}
	return s.conn.WriteMessage(ws.BinaryMessage, message)
}

func (s *Session) writePump() {
	defer s.Close()

	for {
		select {
		case msg := <-s.send:
			var err error
			switch s.server.payloadType {
			case PayloadTypeBinary:
				if err = s.sendBinaryMessage(msg); err != nil {
					LogError("write binary message error: ", err)
					return
				}
				break
			case PayloadTypeText:
				if err = s.SendTextMessage(msg); err != nil {
					LogError("write text message error: ", err)
					return
				}
				break
			}
		}
	}
}

func (s *Session) readPump() {
	defer s.Close()
	for {
		if s.conn == nil {
			break
		}

		messageType, data, err := s.conn.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseNormalClosure, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				LogError("read message error:", err)
			}
			return
		}
		switch messageType {
		case ws.CloseMessage:
			return
		case ws.BinaryMessage:
		case ws.TextMessage:
			_ = s.server.messageHandler(s.SessionId(), data)
			break
		case ws.PingMessage:
			if err = s.sendPongMessage(""); err != nil {
				LogError("write pong message error:", err)
				return
			}
			break
		case ws.PongMessage:
			break
		}
	}
}

func (s *Session) Close() {
	s.server.unregister <- s
	s.closeConnect()
}

func (s *Session) closeConnect() {
	if s.conn != nil {
		if err := s.conn.Close(); err != nil {
			LogError("disconnect error:", err)
		}

		s.conn = nil
	}
}

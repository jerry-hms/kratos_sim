package websocket

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	ws "github.com/gorilla/websocket"
	"google.golang.org/grpc/encoding"
	"net"
	"net/http"
	"time"
)

type Binder func() Any

type ConnectHandler func(SessionID, bool)

type MessageHandler func(SessionID, MessagePayload) error

type HandlerData struct {
	Handler MessageHandler
	Binder  Binder
}

type MessageHandlerMap map[MessageType]*HandlerData

type Server struct {
	*http.Server

	lis      net.Listener
	tlsConf  *tls.Config
	upgrader *ws.Upgrader

	network     string
	address     string
	path        string
	strictSlash bool

	timeout time.Duration

	err   error
	codec encoding.Codec

	messageHandlers MessageHandlerMap

	sessionMgr *SessionManager

	register   chan *Session
	unregister chan *Session

	payloadType PayloadType
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network:     "tcp",
		address:     ":0",
		timeout:     1 * time.Second,
		strictSlash: true,
		path:        "/",

		messageHandlers: make(MessageHandlerMap),

		sessionMgr: NewSessionManager(),

		upgrader: &ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		register:   make(chan *Session),
		unregister: make(chan *Session),

		payloadType: PayloadTypeBinary,
	}

	srv.init(opts...)

	srv.err = srv.listen()

	return srv
}

func (s *Server) Name() string {
	return string(KindWebsocket)
}

func (s *Server) init(opts ...ServerOption) {
	for _, f := range opts {
		f(s)
	}

	s.Server = &http.Server{
		TLSConfig: s.tlsConf,
	}

	http.HandleFunc(s.path, s.wsHandle)
}

func (s *Server) SessionCount() int {
	return s.sessionMgr.Count()
}

func (s *Server) RegisterMessageHandler(messageType MessageType, handler MessageHandler, binder Binder) {
	if _, ok := s.messageHandlers[messageType]; ok {
		return
	}

	s.messageHandlers[messageType] = &HandlerData{
		Handler: handler,
		Binder:  binder,
	}
}

func (s *Server) DeregisterMessageHandler(messageType MessageType) {
	delete(s.messageHandlers, messageType)
}

func RegisterServerMessageHandler[T any](srv *Server, messageType MessageType, handler func(SessionID, *T) error) {
	srv.RegisterMessageHandler(messageType,
		func(id SessionID, payload MessagePayload) error {
			switch t := payload.(type) {
			case *T:
				return handler(id, t)
			default:
				LogError("invalid payload struct type:", t)
				return errors.New("invalid payload struct type")
			}
		},
		func() Any {
			var t T
			return &t
		},
	)
}

func (s *Server) wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)

	if err != nil {
		LogError("upgrade exception:", err)
		return
	}

	session := NewSession(conn, s)
	session.server.register <- session

	session.Listen()
}

func (s *Server) SendMessage(id SessionID, messageType MessageType, message MessagePayload) error {
	var err error
	var msg []byte
	conn, _ := s.sessionMgr.Get(id)
	if conn == nil {
		return errors.New(fmt.Sprintf("%s websocket is no such connection", id))
	}

	msg, err = s.marshalMessage(messageType, message)
	if err != nil {
		return err
	}

	switch s.payloadType {
	case PayloadTypeBinary:
		err = conn.sendBinaryMessage(msg)
	case PayloadTypeText:
		err = conn.SendTextMessage(msg)
	}
	return err
}

func (s *Server) marshalMessage(messageType MessageType, message MessagePayload) ([]byte, error) {
	var err error
	var buf []byte
	switch s.payloadType {
	case PayloadTypeBinary:
		var msg BinaryMessage
		msg.Type = messageType
		msg.Body, err = Marshal(s.codec, message)
		if err != nil {
			return nil, err
		}
		buf, err = msg.Marshal()
		if err != nil {
			return nil, err
		}
	case PayloadTypeText:
		var msg TextMessage
		msg.Type = messageType
		str, err := json.Marshal(message)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(str, &msg.Body)
		if err != nil {
			return nil, err
		}
		buf, err = msg.Marshal()
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func (s *Server) unmarshalMessage(buf []byte) (*HandlerData, MessagePayload, error) {
	var handler *HandlerData
	var payload MessagePayload

	switch s.payloadType {
	case PayloadTypeBinary:
		var msg BinaryMessage
		if err := msg.Unmarshal(buf); err != nil {
			LogError("decode message exception:", err)
			return nil, nil, err
		}
		var ok bool
		handler, ok = s.messageHandlers[msg.Type]
		if !ok {
			LogError("message handler not found:", msg.Type)
			return nil, nil, errors.New("message handler not found")
		}

		if handler.Binder != nil {
			payload = handler.Binder()
		} else {
			payload = msg.Body
		}

		if err := Unmarshal(s.codec, msg.Body, &payload); err != nil {
			LogError("unmarshal message exception:", err)
			return nil, nil, err
		}
	case PayloadTypeText:
		var msg TextMessage
		if err := msg.Unmarshal(buf); err != nil {
			LogError("decode message exception:", err)
			return nil, nil, err
		}

		var ok bool
		handler, ok = s.messageHandlers[msg.Type]
		if !ok {
			LogError("message handler not found:", msg.Type)
			return nil, nil, errors.New("message handler not found")
		}
		if handler.Binder != nil {
			payload = handler.Binder()
		} else {
			payload = msg.Body
		}
		marshal, err := json.Marshal(msg.Body)
		if err != nil {
			return nil, nil, err
		}
		if err := Unmarshal(s.codec, marshal, payload); err != nil {
			LogError("unmarshal message exception:", err)
			return nil, nil, err
		}
	}

	return handler, payload, nil
}

func (s *Server) messageHandler(sessionId SessionID, buf []byte) error {
	var err error
	var handler *HandlerData
	var payload MessagePayload
	if handler, payload, err = s.unmarshalMessage(buf); err != nil {
		LogError("unmarshal message failed :", err)
		return err
	}

	if err = handler.Handler(sessionId, payload); err != nil {
		LogError("message handler failed :", err)
		return err
	}

	return nil
}

func (s *Server) listen() error {
	if s.lis == nil {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return err
		}
		s.lis = lis
	}

	return nil
}

func (s *Server) run() {
	for {
		select {
		case client := <-s.register:
			s.sessionMgr.Add(client)
			s.sessionMgr.CallConnectHandler(client)
		case client := <-s.unregister:
			s.sessionMgr.Remove(client)
			s.sessionMgr.CallConnectHandler(client)
		}
	}
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	LogInfof("server listening on:", s.lis.Addr().String())

	go s.run()

	var err error
	if s.tlsConf != nil {
		err = s.ServeTLS(s.lis, "", "")
	} else {
		err = s.Serve(s.lis)
	}

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	LogInfo("server stopping")
	return s.Shutdown(ctx)
}

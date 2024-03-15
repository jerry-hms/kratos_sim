package websocket

import (
	"sync"
)

type SessionMap map[SessionID]*Session

type SessionManager struct {
	sessions       SessionMap
	mtx            sync.RWMutex
	connectHandler ConnectHandler
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(SessionMap),
	}
}

func (s *SessionManager) Count() int {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return len(s.sessions)
}

func (s *SessionManager) RegisterConnectHandler(handler ConnectHandler) {
	s.connectHandler = handler
}

func (s *SessionManager) Add(c *Session) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.sessions[c.SessionId()] = c
}

func (s *SessionManager) CallConnectHandler(c *Session) {
	//s.mtx.Lock()
	//defer s.mtx.Unlock()

	if s.connectHandler != nil {
		s.connectHandler(c.SessionId(), true)
	}
}

func (s *SessionManager) Remove(c *Session) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for k, v := range s.sessions {
		if c == v {
			delete(s.sessions, k)
			return
		}
	}
}

func (s *SessionManager) Clean() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.sessions = make(SessionMap)
}

func (s *SessionManager) Get(id SessionID) (*Session, bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	sen, ok := s.sessions[id]
	return sen, ok
}

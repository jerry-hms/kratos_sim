package websocket

import (
	"crypto/tls"
	"google.golang.org/grpc/encoding"
	"net"
	"time"
)

type PayloadType uint8

const (
	PayloadTypeBinary = 0
	PayloadTypeText   = 1
	PayloadTypeString = 2
)

type ServerOption func(s *Server)

func WithNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithPath(path string) ServerOption {
	return func(s *Server) {
		s.path = path
	}
}

func WithConnectHandle(h ConnectHandler) ServerOption {
	return func(s *Server) {
		s.sessionMgr.RegisterConnectHandler(h)
	}
}

func WithTLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
		s.tlsConf = c
	}
}

func WithListener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.codec = encoding.GetCodec(c)
	}
}

func WithChannelBufferSize(size int) ServerOption {
	return func(_ *Server) {
		channelBufSize = size
	}
}

func WithPayloadType(payloadType PayloadType) ServerOption {
	return func(s *Server) {
		s.payloadType = payloadType
	}
}

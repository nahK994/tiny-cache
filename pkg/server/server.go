package server

import (
	"log/slog"
	"net"
)

type Server struct {
	listenAddress string
	ln            net.Listener
}

func InitiateServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn.RemoteAddr().String(), conn)
		go peer.readConn()
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	s.ln = ln
	defer ln.Close()

	slog.Info("server running", "listenAddr", s.listenAddress)

	return s.acceptConn()
}

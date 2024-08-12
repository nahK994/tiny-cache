package server

import (
	"log/slog"
	"net"
)

type Config struct {
	ListenAddress string
}

type Server struct {
	Config
	peers map[*Peer]bool
	ln    net.Listener
}

func NewServer(cfg Config) *Server {
	return &Server{
		Config: cfg,
		peers:  make(map[*Peer]bool),
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn)
		s.peers[peer] = true
		go peer.readConn()
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return err
	}

	s.ln = ln
	defer ln.Close()

	slog.Info("server running", "listenAddr", s.ListenAddress)

	return s.acceptConn()
}

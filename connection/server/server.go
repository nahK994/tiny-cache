package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/nahK994/TinyCache/pkg/config"
)

type Server struct {
	listenAddress string
	ln            net.Listener
}

func InitiateServer() *Server {
	return &Server{
		listenAddress: fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := newPeer(conn.RemoteAddr().String(), conn)
		go peer.handleConn()
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

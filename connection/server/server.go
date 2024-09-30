package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/nahK994/TinyCache/connection/server/handlers"
	"github.com/nahK994/TinyCache/pkg/config"
)

type Server struct {
	listenAddress string
	ln            net.Listener
}

type Peer struct {
	clientAddr string
	conn       net.Conn
}

func Init() *Server {
	return &Server{
		listenAddress: fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
	}
}

func newPeer(addr string, conn net.Conn) *Peer {
	return &Peer{
		clientAddr: addr,
		conn:       conn,
	}
}

func processAsyncTasks() {
	c := config.App.Cache
	for {
		select {
		case <-config.App.FlushCh:
			c.FLUSHALL()
		}
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := newPeer(conn.RemoteAddr().String(), conn)

		go processAsyncTasks()
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

func (p *Peer) handleConn() {
	fmt.Printf("\nPaired with %s\n\n", p.clientAddr)
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("peer read error", "err", err, "client", p.clientAddr)
			p.conn.Close()
			return
		}

		rawCmd := string(buf[:n])
		formattedCmd := ""
		for _, ch := range rawCmd {
			if ch == '\r' {
				formattedCmd += "\\r"
			} else if ch == '\n' {
				formattedCmd += "\\n"
			} else {
				formattedCmd += string(ch)
			}
		}
		fmt.Printf("%s> %s\n", p.clientAddr, formattedCmd)

		var res string
		if output, err := handlers.HandleCommand(rawCmd); err != nil {
			res = err.Error()
		} else {
			res = output
		}
		p.conn.Write([]byte(res))
	}
}

package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/nahK994/TinyCache/pkg/handlers"
)

type Peer struct {
	clientAddr string
	conn       net.Conn
}

func newPeer(addr string, conn net.Conn) *Peer {
	return &Peer{
		clientAddr: addr,
		conn:       conn,
	}
}

func (p *Peer) handleConn() {
	slog.Info("Paired with", "client", p.clientAddr)
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("peer read error", "err", err, "client", p.clientAddr)
			p.conn.Close()
			return
		}

		fmt.Printf("%s> %s", p.clientAddr, string(buf[:n]))
		resp, err := handlers.HandleCommand(string(buf[:n]))
		if err != nil {
			p.conn.Write([]byte(err.Error()))
		} else {
			p.conn.Write([]byte(resp))
		}
	}
}

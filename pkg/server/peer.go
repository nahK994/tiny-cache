package server

import (
	"log/slog"
	"net"
)

type Peer struct {
	clientAddr string
	conn       net.Conn
}

func NewPeer(addr string, conn net.Conn) *Peer {
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

		slog.Info("request from client", p.clientAddr, string(buf[:n-1]))
		p.conn.Write([]byte("+OK\r\n"))
	}
}

package server

import (
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) readConn() {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("peer read error", "err", err, "remoteAddr", p.conn.RemoteAddr())
			p.conn.Close()
			return
		}

		fmt.Println(string(buf[:n]))
		p.conn.Write([]byte("+Ok\r\n"))
	}
}

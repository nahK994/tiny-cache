package server

import (
	"net"

	"github.com/nahK994/ScratchCache/handlers"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) readConn() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])

		p.conn.Write([]byte("+OK\r\n"))

		// p.msgCh <- msgBuf
		handlers.HandleCommand(msgBuf)
	}
}

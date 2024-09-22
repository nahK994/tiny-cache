package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/nahK994/TinyCache/pkg/handlers"
	"github.com/nahK994/TinyCache/pkg/utils"
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
		if err := utils.ValidateSerializedCmd(rawCmd); err != nil {
			res = err.Error()
		} else {
			res = handlers.HandleCommand(rawCmd)
		}
		p.conn.Write([]byte(res))
	}
}

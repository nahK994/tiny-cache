package client

import (
	"bufio"
	"log/slog"
	"net"
	"os"
)

type Client struct {
	dialingAddr string
}

func InitClient(addr string) *Client {
	return &Client{
		dialingAddr: addr,
	}
}

func (c *Client) Start() error {
	conn, err := net.Dial("tcp", c.dialingAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	slog.Info("Paired with", "server", c.dialingAddr)
	userReader := bufio.NewReader(os.Stdin)

	for {
		str, _ := userReader.ReadString('\n')
		conn.Write([]byte(str))
	}
}

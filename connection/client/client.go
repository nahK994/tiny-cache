package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

type Client struct {
	dialingAddr string
	conn        net.Conn
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
	c.conn = conn
	defer conn.Close()

	slog.Info("Paired with", "server", c.dialingAddr)
	return c.handleConn()
}

func (c *Client) handleConn() error {
	buf := make([]byte, 1024)
	userReader := bufio.NewReader(os.Stdin)

	for {
		str, _ := userReader.ReadString('\n')
		c.conn.Write([]byte(str))

		n, err := c.conn.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println(string(buf[:n]))
	}
}

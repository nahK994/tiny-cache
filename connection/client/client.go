package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/nahK994/TinyCache/pkg/resp"
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

	// slog.Info("Paired with", "server", c.dialingAddr)
	fmt.Printf("Paired with server %s\n\n", c.dialingAddr)
	return c.handleConn()
}

func (c *Client) handleConn() error {
	buf := make([]byte, 1024)
	userReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("(%s) client-cli> ", c.conn.LocalAddr())
		str, _ := userReader.ReadString('\n')

		serializedCmd, err := resp.Serialize(str[:len(str)-1])
		var resp string
		if err != nil {
			resp = err.Error()
		} else {
			resp = serializedCmd
		}
		c.conn.Write([]byte(resp))

		n, err := c.conn.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println(string(buf[:n]))
	}
}

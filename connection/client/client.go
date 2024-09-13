package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
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
	fmt.Printf("Paired with server %s\n", c.dialingAddr)
	fmt.Printf("%s\n\n", utils.GetClientMessage())
	return c.handleConn()
}

func (c *Client) handleConn() error {
	buf := make([]byte, 1024)
	userReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("(%s) client-cli> ", c.conn.LocalAddr())
		str, _ := userReader.ReadString('\n')

		str = str[:len(str)-1] // skip last \n
		var response string
		if err := utils.ValidateRawCommand(str); err != nil {
			response = err.Error()
		} else {
			response = resp.Serialize(str)
		}
		c.conn.Write([]byte(response))

		n, err := c.conn.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println(string(buf[:n]))
	}
}

package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

type Client struct {
	dialingAddr string
	conn        net.Conn
}

func InitClient() *Client {
	return &Client{
		dialingAddr: fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
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
		// fmt.Printf("(%s) client-cli> ", c.conn.LocalAddr())
		fmt.Printf("tinycache-cli> ")
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

		deserializedResp := resp.Deserializer(string(buf[:n]))
		res := processResp(deserializedResp)
		fmt.Println(res)
	}
}

func processResp(res interface{}) string {
	switch v := res.(type) {
	case int:
		return fmt.Sprintln("(integer)", v)
	case string:
		if len(v) == 0 {
			return fmt.Sprintln("(nil)")
		} else {
			return fmt.Sprintln(v)
		}
	case []string:
		var res string
		for i, item := range v {
			res += fmt.Sprintf("%d) %s\n", i+1, item)
		}
		return res
	case error:
		return fmt.Sprintln("(error)", v.Error())
	default:
		return ""
	}
}

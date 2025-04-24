package client

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/nahK994/tiny-cache/pkg/config"
	"github.com/nahK994/tiny-cache/pkg/resp"
	"github.com/nahK994/tiny-cache/pkg/validators"
)

type Client struct {
	dialingAddr string
	conn        net.Conn
	rl          *readline.Instance // To handle arrow keys and line history
}

func Init() *Client {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:            "tinycache-cli> ",
		HistoryFile:       "/tmp/readline.tmp", // For persistent history
		InterruptPrompt:   "^C",
		HistorySearchFold: true,
	})
	if err != nil {
		fmt.Println("Failed to initialize readline:", err)
		os.Exit(1)
	}

	return &Client{
		dialingAddr: fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
		rl:          rl,
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

	var clientMessage string = fmt.Sprintf(
		"\nPlease use these following commands:\n%s\n%s\n%s\n%s\n\nType ^C to exit...\n",
		strings.Join([]string{
			resp.PING, resp.SET, resp.GET, resp.EXISTS,
		}, ", "),
		strings.Join([]string{
			resp.FLUSHALL, resp.DEL, resp.INCR, resp.DECR,
		}, ", "),
		strings.Join([]string{
			resp.LPUSH, resp.LPOP, resp.LRANGE, resp.RPUSH, resp.RPOP,
		}, ", "),
		strings.Join([]string{
			resp.EXPIRE, resp.TTL, resp.PERSIST,
		}, ", "),
	)
	fmt.Printf("%s\n\n", clientMessage)
	return c.handleConn()
}

func (c *Client) handleConn() error {
	buf := make([]byte, 1024)

	for {
		str, err := c.rl.Readline()
		if err == readline.ErrInterrupt {
			if len(str) == 0 {
				break // Exit on Ctrl+C
			} else {
				continue
			}
		}

		if err := validators.ValidateRawCommand(str); err != nil {
			c.conn.Write([]byte(err.Error()))
		} else {
			c.conn.Write([]byte(resp.Serialize(str)))
		}

		n, err := c.conn.Read(buf)
		if err != nil {
			return err
		}

		deserializedResp := resp.Deserializer(string(buf[:n]))
		res := processResp(deserializedResp)
		fmt.Println(res)
	}

	return nil
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

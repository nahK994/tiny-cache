package errors

import "fmt"

type Err struct {
	Msg string
}

func (e Err) Error() string {
	return e.Msg
}

type InvalidCmd struct {
	Cmd string
}

func (e InvalidCmd) Error() string {
	return fmt.Sprintf("-ERR unknown command '%s'\r\n", e.Cmd)
}

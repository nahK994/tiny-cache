package errors

import "fmt"

type Err struct {
	Msg  string
	File string
	Line int
}

func (e Err) Error() string {
	return e.Msg
}

func (e Err) ErrorInfo() string {
	return fmt.Sprintln("error found in", e.File, "from", e.Line)
}

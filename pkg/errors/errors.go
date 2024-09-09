package errors

type MalformedErr struct {
	Msg string
}

func (e MalformedErr) Error() string {
	return e.Msg
}

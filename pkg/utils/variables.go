package utils

type respSynt struct {
	STRING  rune
	ERROR   rune
	INTEGER rune
	BULK    rune
	ARRAY   rune
}

var (
	RespSyntax respSynt = respSynt{
		STRING:  '+',
		ERROR:   '-',
		INTEGER: ':',
		BULK:    '$',
		ARRAY:   '*',
	}
)

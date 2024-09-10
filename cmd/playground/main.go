package main

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/resp"
)

func main() {
	a := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$12\r\nvalue\r\nvalue\r\n"
	str, err := resp.Deserializer(a)
	fmt.Println(str)
	fmt.Println(err)

	// c := '\r'
	// d := '\n'
	// fmt.Println(c, d)
}

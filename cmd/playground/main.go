package main

import (
	"fmt"

	"github.com/nahK994/SimpleCache/pkg/resp"
)

func test(cmd string) {
	str, _ := resp.Deserializer(cmd)
	fmt.Println(cmd, str)
}

func main() {
	test("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
	test("*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n")
	test("*3\r\n$6\r\nEXPIRE\r\n$3\r\nkey\r\n$2\r\n10\r\n")
	test("*2\r\n$3\r\nDEL\r\n$3\r\nkey\r\n")
	test("*2\r\n$4\r\nINCR\r\n$3\r\nkey\r\n")
	test("*3\r\n$5\r\nLPUSH\r\n$4\r\nlist\r\n$5\r\nvalue\r\n")
	test("*4\r\n$6\r\nLRANGE\r\n$4\r\nlist\r\n$1\r\n0\r\n$2\r\n10\r\n")
	test("*2\r\n$4\r\nAUTH\r\n$5\r\nmyPwd\r\n")
	test("*1\r\n$4\r\nPING\r\n")
	test("*6\r\n$5\r\nHMSET\r\n$4\r\nhash\r\n$6\r\nfield1\r\n$6\r\nvalue1\r\n$6\r\nfield2\r\n$6\r\nvalue2\r\n")
}

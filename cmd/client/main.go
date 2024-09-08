package main

import (
	"fmt"
	"strings"
)

func main() {
	setCommand := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	fmt.Println(strings.Split(setCommand, "\r\n"))
}

// setCommand := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
// fmt.Println(setCommand)
// getCommand := "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"
// fmt.Println(getCommand)
// expireCommand := "*3\r\n$6\r\nEXPIRE\r\n$3\r\nkey\r\n$2\r\n10\r\n"
// fmt.Println(expireCommand)
// delCommand := "*2\r\n$3\r\nDEL\r\n$3\r\nkey\r\n"
// fmt.Println(delCommand)
// incrCommand := "*2\r\n$4\r\nINCR\r\n$3\r\nkey\r\n"
// fmt.Println(incrCommand)
// lpushCommand := "*3\r\n$5\r\nLPUSH\r\n$4\r\nlist\r\n$5\r\nvalue\r\n"
// fmt.Println(lpushCommand)
// lrangeCommand := "*4\r\n$6\r\nLRANGE\r\n$4\r\nlist\r\n$1\r\n0\r\n$2\r\n10\r\n"
// fmt.Println(lrangeCommand)
// authCommand := "*2\r\n$4\r\nAUTH\r\n$5\r\nmyPwd\r\n"
// fmt.Println(authCommand)
// pingCommand := "*1\r\n$4\r\nPING\r\n"
// fmt.Println(pingCommand)
// hmsetCommand := "*6\r\n$5\r\nHMSET\r\n$4\r\nhash\r\n$5\r\nfield1\r\n$5\r\nvalue1\r\n$5\r\nfield2\r\n$5\r\nvalue2\r\n"
// fmt.Println(hmsetCommand)

// // 1. LPUSH to initialize the list with value "1"
// lpushInitialCommand := "*3\r\n$5\r\nLPUSH\r\n$3\r\nkey\r\n$1\r\n1\r\n"
// fmt.Println("Initialize list with '1':\n", lpushInitialCommand)

// // 2. RPUSH to add value "2" to the list
// rpushSecondCommand := "*3\r\n$5\r\nRPUSH\r\n$3\r\nkey\r\n$1\r\n2\r\n"
// fmt.Println("Push '2' to the list:\n", rpushSecondCommand)

// // 3. RPUSH to add value "3" to the list
// rpushThirdCommand := "*3\r\n$5\r\nRPUSH\r\n$3\r\nkey\r\n$1\r\n3\r\n"
// fmt.Println("Push '3' to the list:\n", rpushThirdCommand)

// // 4. LRANGE to retrieve values between index 0 and 2
// lrangeCommand := "*4\r\n$6\r\nLRANGE\r\n$3\r\nkey\r\n$1\r\n0\r\n$1\r\n2\r\n"
// fmt.Println("Get values between 0 and 2:\n", lrangeCommand)

package main

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func main() {
	c := cache.InitCache()
	// handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\ncounter\r\n$2\r\n10\r\n")
	c.SET("counter", "10")
	fmt.Println("get =", c.GET("counter"))
	// response, _ := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$7\r\ncounter\r\n")
	// fmt.Println("response =", response)
	// fmt.Println("INCR =", c.INCR("counter"))
}

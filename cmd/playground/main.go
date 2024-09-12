package main

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/resp"
)

func main() {
	str := "SET name Shomi Khan"
	fmt.Println(resp.Serialize(str))
}

package main

import (
	"log"

	"github.com/nahK994/SimpleCache/pkg/client"
)

func main() {
	client := client.InitClient("127.0.0.1:8000")
	log.Fatal(client.Start())
}

package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection"
)

func main() {
	client := connection.InitClient()
	log.Fatal(client.Start())
}

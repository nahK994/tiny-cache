package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection/client"
)

func main() {
	client := client.InitClient()
	log.Fatal(client.Start())
}

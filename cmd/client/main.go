package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection/client"
)

func main() {
	client := client.Init()
	log.Fatal(client.Start())
}

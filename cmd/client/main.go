package main

import (
	"log"

	"github.com/nahK994/tiny-cache/connection/client"
)

func main() {
	client := client.Init()
	log.Fatal(client.Start())
}

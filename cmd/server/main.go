package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection"
)

func main() {
	srv := connection.InitiateServer()
	log.Fatal(srv.Start())
}

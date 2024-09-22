package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection/server"
)

func main() {
	srv := server.InitiateServer()
	log.Fatal(srv.Start())
}

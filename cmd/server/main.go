package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection/server"
)

func main() {
	srv := server.InitiateServer("127.0.0.1:8000")
	log.Fatal(srv.Start())
}

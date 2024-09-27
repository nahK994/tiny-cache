package main

import (
	"log"

	"github.com/nahK994/TinyCache/connection/server"
)

func main() {
	srv := server.Init()
	log.Fatal(srv.Start())
}

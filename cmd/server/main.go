package main

import (
	"log"

	"github.com/nahK994/tiny-cache/connection/server"
)

func main() {
	srv := server.Init()
	log.Fatal(srv.Start())
}

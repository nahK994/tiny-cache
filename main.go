package main

import (
	"log"

	"github.com/nahK994/ScratchCache/server"
)

func main() {
	srv := server.NewServer(server.Config{
		ListenAddress: "127.0.0.1:8000",
	})
	log.Fatal(srv.Start())
}

package main

import (
	"log"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/server"
	"github.com/nahK994/TinyCache/handlers"
)

func main() {
	srv := server.InitiateResp("127.0.0.1:8000")
	srv.RequestHandler(func(request models.RespRequest, response *models.RespResponse) {
		handlers.HandleCommand(request.Request)
		response.Response = "+OK\r\n"
	})
	log.Fatal(srv.Start())
}

package main

import (
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/sjug/am-go/database"
	"github.com/sjug/am-go/server"
)

func main() {
	const port = ":8080"
	httpRouter := httptreemux.New()
	server.InitDetails(httpRouter)
	server.InitTier(httpRouter)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, httpRouter))
}

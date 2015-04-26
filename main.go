package main

import (
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/sjug/am-go/server"
)

func main() {
	const port = ":8080"
	router := httptreemux.New()
	server.InitDetails(router)
	//server.InitTier(router)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

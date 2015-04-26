package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/sjug/am-go/server"
)

func userTierHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// TODO userTierHandler
	fmt.Fprintf(w, "hello, %s!\n", ps["num"])
}

func main() {
	const port = ":8080"
	router := httptreemux.New()
	server.InitDetails(router)
	router.GET("/user/tier/:num", userTierHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/dimfeld/httptreemux"
	"log"
	"net/http"
	"strconv"
)

type CollectorDetails struct {
	CollectorNumber int    `json:"number"`
	CollectorName   string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	num, err := strconv.Atoi(ps["num"])
	if err != nil || num < 1 {
		http.Error(w, "Please enter a numeric collector number.", http.StatusInternalServerError)
		return
	}
	resp := &CollectorDetails{
		CollectorNumber: num,
		CollectorName:   "George"}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Responded with ", string(json))
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func userTierHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// TODO userTierHandler
	fmt.Fprintf(w, "hello, %s!\n", ps["num"])
}

func main() {
	const port = ":8080"
	router := httptreemux.New()
	router.GET("/user/:num", userHandler)
	router.GET("/user/tier/:num", userTierHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimfeld/httptreemux"
	"github.com/sjug/am-go/database"
	"github.com/sjug/am-go/structure"
)

func userHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	num, err := strconv.Atoi(ps["num"])
	if err != nil || num < 1 {
		http.Error(w, "Please enter a numeric collector number.", http.StatusInternalServerError)
		return
	}
	resp, _ := database.GetUserDetailsFromNumber(num)
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Responded with ", string(json))
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// InitDetails intializes routing for details path
func InitDetails(router *httptreemux.TreeMux) {
	// Collector details
	router.GET("/user/:num", userHandler)
}

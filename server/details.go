package details

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimfeld/httptreemux"
)

// CollectorDetails type holds collector data
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

// InitDetails intializes routing for details path
func InitDetails(router *httptreemux.TreeMux) {
	// Collector details
	router.GET("/user/:num", userHandler)
}

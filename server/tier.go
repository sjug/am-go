package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimfeld/httptreemux"
	"github.com/sjug/am-go/soap"
)

func userTierHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	num, err := strconv.Atoi(ps["num"])
	if err != nil || num < 1 {
		http.Error(w, "Please enter a valid collector number.", http.StatusInternalServerError)
		return
	}
	resp, err := soap.GetTierFromSoap(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Responded with ", string(json))
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// InitTier func sets up routing for tier path
func InitTier(router *httptreemux.TreeMux) {
	router.GET("/collector/tier/:num", userTierHandler)
}

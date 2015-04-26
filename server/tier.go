package server

import (
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

func userTierHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// TODO userTierHandler
	fmt.Fprintf(w, "hello, %s!\n", ps["num"])
}

// InitTier func sets up routing for tier path
func InitTier(router *httptreemux.TreeMux) {
	router.GET("/user/tier/:num", userTierHandler)
}

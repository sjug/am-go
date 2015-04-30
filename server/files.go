package server

import (
	"net/http"
	"path/filepath"

	"github.com/dimfeld/httptreemux"
)

func pageHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	http.ServeFile(w, r, filepath.Join("content/", ps["file"]))
	return
}

func fileHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	http.ServeFile(w, r, filepath.Join("content/", ps["filepath"]))
	return
}

// InitStatic sets up routing for static webpages
func InitStatic(router *httptreemux.TreeMux) {
	router.GET("/", pageHandler)
	router.GET("/*filepath", fileHandler)

}

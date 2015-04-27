package server

import (
	"net/http"
	"path/filepath"

	"github.com/dimfeld/httptreemux"
)

func pageHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	//file := "content/" + ps["page"]
	//t, _ := template.ParseFiles(file)
	//t.Execute(w, nil)
	http.ServeFile(w, r, filepath.Join("content/", ps["file"]))
	return
}

func staticHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	http.ServeFile(w, r, filepath.Join("content/", ps["file"]))
	return
}

// InitStatic sets up routing for static webpages
func InitStatic(router *httptreemux.TreeMux) {
	router.GET("/", pageHandler)
	router.GET("/:file", staticHandler)

}

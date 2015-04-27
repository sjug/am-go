package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

func pageHandler(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	file := "content/" + ps["page"]
	fmt.Println(file)
	t, _ := template.New(file).Delims("<<", ">>").ParseFiles("content/index.html")
	t.Execute(w, nil)
}

// InitStatic sets up routing for static webpages
func InitStatic(router *httptreemux.TreeMux) {
	router.GET("/:page", pageHandler)
}

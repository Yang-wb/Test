package main

import (
	"leango/crawler/forntend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/fronted/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe("8888", nil)
	if err != nil {
		panic(err)
	}
}

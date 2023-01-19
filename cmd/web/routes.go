package main

import "net/http"

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (a *application) routes() http.Handler {
	standardMiddleWare := alice.New(a.recoverPanic, a.logRequest, secureHeaders)
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleWare.Then(mux)
}

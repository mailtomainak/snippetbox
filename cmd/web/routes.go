package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (a *application) routes() http.Handler {
	standardMiddleWare := alice.New(a.recoverPanic, a.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(a.home))
	mux.Get("/snippet/create", http.HandlerFunc(a.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(a.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(a.showSnippet))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleWare.Then(mux)
}

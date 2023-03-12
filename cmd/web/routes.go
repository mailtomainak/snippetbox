package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (a *application) routes() http.Handler {

	mux := pat.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	dynamic := alice.New(a.sessionManager.LoadAndSave)

	mux.Get("/", dynamic.ThenFunc(a.home))

	mux.Get("/snippet/create", dynamic.ThenFunc(a.createSnippetForm))
	mux.Post("/snippet/create", dynamic.ThenFunc(a.createSnippet))
	mux.Get("/snippet/:id", dynamic.ThenFunc(a.showSnippet))

	mux.Get("/user/signup", dynamic.ThenFunc(a.userSignup))
	mux.Post("/user/signup", dynamic.ThenFunc(a.userSignupPost))
	mux.Get("/user/login", dynamic.ThenFunc(a.userLogin))
	mux.Post("/user/login", dynamic.ThenFunc(a.userLoginPost))
	mux.Post("/user/logout", dynamic.ThenFunc(a.userLogoutPost))

	standardMiddleWare := alice.New(a.recoverPanic, a.logRequest, secureHeaders)
	return standardMiddleWare.Then(mux)
}

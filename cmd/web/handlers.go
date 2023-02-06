package main

import (
	"fmt"
	"github.com/mailtomainak/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	snippets, err := a.snippetModel.Latest()
	if err != nil {
		a.serverError(w, err)
	}

	a.render(w, r, "home.page.tmpl", &templateData{Snippets: snippets})

}

func (a *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	var s *models.Snippet
	if s, err = a.snippetModel.Get(id); err != nil {
		if err == models.ErrNoRecord {
			a.notFound(w)
			return
		}
		a.serverError(w, err)
		return
	}
	a.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := a.snippetModel.Insert(title, content, expires)

	if err != nil {
		a.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

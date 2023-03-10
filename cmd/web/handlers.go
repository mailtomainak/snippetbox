package main

import (
	"fmt"
	"github.com/mailtomainak/snippetbox/pkg/forms"
	"github.com/mailtomainak/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := a.snippetModel.Latest()
	if err != nil {
		a.serverError(w, err)
	}

	a.render(w, r, "home.page.tmpl", &templateData{Snippets: snippets})

}

func (a *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	flash := a.sessionManager.PopString(r.Context(), "flash")

	a.render(w, r, "show.page.tmpl", &templateData{Snippet: s, Flash: flash})
}

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.NewForm(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		a.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := a.snippetModel.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))

	if err != nil {
		a.serverError(w, err)
		return
	}
	a.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (a *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.NewForm(nil),
	})
}

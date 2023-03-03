package main

import (
	"fmt"
	"github.com/mailtomainak/snippetbox/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
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
	a.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	// Check that the title field is not blank and is not more than 100 characters
	// long. If it fails either of those checks, add a message to the errors
	// map using the field name as the key.

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	// Check that the Content field isn't blank.

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	// Check the expires field isn't blank and matches one of the permitted
	// values ("1", "7" or "365").

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "1" && expires != "7" && expires != "365" {
		errors["title"] = "This field is invalid"
	}

	if len(errors) > 0 {
		a.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := a.snippetModel.Insert(title, content, expires)

	if err != nil {
		a.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/?id=%d", id), http.StatusSeeOther)
}

func (a *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "create.page.tmpl", nil)
	//w.Write([]byte("Create a new snippet..."))
}

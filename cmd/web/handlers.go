package main

import (
	"fmt"
	"github.com/mailtomainak/snippetbox/pkg/models"
	"html/template"
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

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}
	/*files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.errorLog.Println(err.Error())
		a.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		a.errorLog.Println(err.Error())
		a.serverError(w, err)
	}
	*/

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
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}
	err = ts.Execute(w, s)
	if err != nil {
		a.serverError(w, err)
	}
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

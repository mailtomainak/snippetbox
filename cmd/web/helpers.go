package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (a *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func (a *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *application) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}

func (a *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		Flash:       a.sessionManager.PopString(r.Context(), "flash"),
		CurrentYear: time.Now().Year(),
	}
}

func (a *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := a.templateCache[name]
	if !ok {
		a.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, td)
	if err != nil {
		a.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

package main

import (
	"html/template"
	"net/http"
)

func (a App) Welcome(w http.ResponseWriter, r *http.Request) {
	var data map[string]bool

	t, err := template.ParseFiles(GetTemplates("base", "upload"))
	a.Log(err, "Parsing template files")
	a.Log(t.ExecuteTemplate(w, "base", data), "Execute Welcome template")
}
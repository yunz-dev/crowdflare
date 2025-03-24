package handlers

import (
	"html/template"
	"net/http"
)

func AppPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/templates/pages/app.html"))
    tmpl.Execute(w, nil)
}

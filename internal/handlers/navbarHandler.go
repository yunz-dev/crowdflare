package handlers

import (
	"html/template"
	"net/http"
)

func NavBar(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/partials/navbar.html"))
	tmpl.Execute(w, nil)
}

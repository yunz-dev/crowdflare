package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/templates/landingPage.html"))
    tmpl.Execute(w, nil)
    fmt.Fprintf(w, "Server is up and running!") // Response message
}

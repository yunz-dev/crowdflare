package handlers

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
    // Parse the templates (header, footer, and landing page)
    partial := "web/templates/partials/"
    tmpl, err := template.ParseFiles(
    partial+ "navbar.html",
    partial+ "footer.html",
    partial+ "hero.html",
    partial+ "action.html",
    partial+ "guide.html",
    "web/templates/index.html",
    )
    if err != nil {
        log.Fatal(err)
    }

    // Render the template with the data (you can pass any data you want to render here)
    err = tmpl.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        log.Fatal(err)
    }
}

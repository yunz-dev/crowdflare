package handlers

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
    // Parse the templates (header, footer, and landing page)
    tmpl, err := template.ParseFiles("web/templates/partials/navbar.html","web/templates/partials/footer.html", "web/templates/index.html")
    if err != nil {
        log.Fatal(err)
    }

    // Render the template with the data (you can pass any data you want to render here)
    err = tmpl.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        log.Fatal(err)
    }
}

package handlers

import (
	"html/template"
	"net/http"
	"github.com/yunz-dev/crowdflare/internal/models"
)
func LeaderboardPage(w http.ResponseWriter, r *http.Request) {

	data := models.LeaderboardData{
		TopUsers: []models.User{
			{Username: "John Doe", Karma: 100},
			{Username: "Jane Karma", Karma: 90},
			{Username: "John Doe", Karma: 73},
		},
		Contributors: []models.User{
			{Username: "Queen Elizabeth", Karma: 53},
			{Username: "Prince Henry", Karma: 32},
			{Username: "Bob Jane", Karma: 23},
		},
	}
	
    tmpl := template.Must(template.ParseFiles(
		"web/templates/leaderboardPage.html",
		"web/templates/partials/leaderboardData.html",
	))
    tmpl.Execute(w, data)
}
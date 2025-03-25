package handlers

import (
	"html/template"
	"net/http"
)
func LeaderboardPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/templates/leaderboardPage.html"))
    tmpl.Execute(w, nil)
}

func LeaderboardData(w http.ResponseWriter, r *http.Request) {
	timeframe := r.URL.Path[len("/leaderboard/data/"):] // Extract "all-time", "month", "week"
    
    if timeframe != "all-time" && timeframe != "month" && timeframe != "week" {
        http.Error(w, "Invalid timeframe. Use /leaderboard/all-time, /leaderboard/month, or /leaderboard/week", http.StatusBadRequest)
        return
    }
	tmpl := template.Must(template.ParseFiles("web/templates/partials/leaderboardData.html"))
	tmpl.Execute(w, nil)
}
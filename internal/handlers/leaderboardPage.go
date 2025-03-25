package handlers

import (
	"html/template"
	"net/http"
	"github.com/yunz-dev/crowdflare/internal/models"
	"github.com/yunz-dev/crowdflare/internal/services"
)
func LeaderboardPage(w http.ResponseWriter, r *http.Request) {

    tmpl := template.Must(template.ParseFiles("web/templates/leaderboardPage.html"))
	tmpl.Execute(w, nil)

}

func LeaderboardData(w http.ResponseWriter, r *http.Request) {

    timeframe := r.URL.Query().Get("timeframe")
    if timeframe == "" {
        timeframe = "all" // Default value
    }
	data := models.LeaderboardData{
		CurrentUser: models.User{Username: ""},
		Users: []models.User{
			{
				Username: "John Doe",
				Karma: 100,
				KarmaHistory: []int{100, 102, 105, 108, 110, 112, 115, 117, 119, 121, 123, 124, 126, 128, 130, 132, 135, 137, 140, 142, 145, 147, 150, 153, 155, 158, 160, 163, 165, 167},
			},
			{
				Username: "Jane Karma",
				Karma: 90,
				KarmaHistory: []int{90, 92, 95, 97, 100, 102, 105, 107, 110, 112, 114, 116, 119, 121, 123, 125, 128, 130, 132, 134, 136, 139, 141, 143, 145, 147, 149, 151, 153, 155},
			},
			{
				Username: "Queen Elizabeth",
				Karma: 53,
				KarmaHistory: []int{53, 55, 58, 60, 62, 64, 67, 69, 71, 74, 76, 78, 80, 82, 85, 87, 89, 92, 94, 97, 99, 101, 103, 106, 108, 110, 112, 114, 116, 118},
			},
			{
				Username: "Prince Henry",
				Karma: 32,
				KarmaHistory: []int{32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90},
			},
			{
				Username: "Bob Jane",
				Karma: 23,
				KarmaHistory: []int{23, 25, 27, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82},
			},
			{
				Username: "John Doe",
				Karma: 73,
				KarmaHistory: []int{73, 75, 77, 80, 82, 84, 87, 89, 91, 93, 95, 97, 100, 102, 104, 106, 109, 111, 113, 115, 117, 119, 121, 123, 125, 128, 130, 132, 134, 136},
			},
		},
	}
	service.SliceDataByTime(&data, timeframe)
    tmpl := template.Must(template.ParseFiles("web/templates/partials/leaderboardData.html"))
    tmpl.Execute(w, data)
}
package handlers

import (
	"html/template"
	"net/http"
	"github.com/yunz-dev/crowdflare/internal/services"
	"strconv"
)

func NearbyBuildings(w http.ResponseWriter, r *http.Request) {
	// Get the user's current location
	latPar := r.URL.Query().Get("lat")
	longPar := r.URL.Query().Get("long")
	// date := r.URL.Query().Get("date")

	lat64, err := strconv.ParseFloat(latPar, 32)
    if err != nil {
        http.Error(w, "Invalid latitude", http.StatusBadRequest)
        return
    }
    
    long64, err := strconv.ParseFloat(longPar, 32)
    if err != nil {
        http.Error(w, "Invalid longitude", http.StatusBadRequest)
        return
    }
	lat := float32(lat64)
	long := float32(long64)

	// Get the nearby rooms
	data := service.GetFullData()
	rooms := service.GetNearbyBuildings(lat, long, 0.0045, data)

	// Render the template
	tmpl := template.Must(template.ParseFiles("web/templates/nearbyRooms.html"))
	tmpl.Execute(w, rooms)
}
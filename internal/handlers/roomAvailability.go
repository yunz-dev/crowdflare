package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/yunz-dev/crowdflare/internal/services"
	"strconv"
)

func NearbyBuildings(w http.ResponseWriter, r *http.Request) {
	// Get the user's current location
	latPar := r.URL.Query().Get("lat")
	longPar := r.URL.Query().Get("lng")

	lat64, err := strconv.ParseFloat(latPar, 64)
    if err != nil {
        http.Error(w, "Invalid latitude", http.StatusBadRequest)
        return
    }
    
    long64, err := strconv.ParseFloat(longPar, 64)
    if err != nil {
        http.Error(w, "Invalid longitude", http.StatusBadRequest)
        return
    }
	lat := float32(lat64)
	long := float32(long64)

	// Get the nearby rooms
	data := service.GetFullData()
	buildings := service.GetNearbyBuildings(lat, long, 0.0045, data)
	freeBuildings := service.GetFreeBuildings(buildings)

	fmt.Println(freeBuildings)
	// Render the template
	tmpl := template.Must(template.ParseFiles("web/templates/nearbyRooms.html"))
	tmpl.Execute(w, freeBuildings)
}
package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"time"
	"math"
)

func GetFullData() map[string]interface{} {
	// Define the GraphQL query
	query := `{"query": "query MyQuery { buildings { name long lat id rooms { bookings { end start name } } } }"}`
	reqBody := []byte(query)

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://graphql.csesoc.app/v1/graphql", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	// Parse the response body into a map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}

	// Navigate to buildings --> rooms --> bookings
	data, exists := result["data"].(map[string]interface{})
	if !exists {
		return nil
	}

	buildings, exists := data["buildings"].([]interface{})
	if !exists {
		return nil
	}

	for _, b := range buildings {
		building := b.(map[string]interface{})
		rooms, exists := building["rooms"].([]interface{})
		if !exists {
			continue // Skip if no rooms
		}

		for _, r := range rooms {
			room := r.(map[string]interface{})
			bookings, exists := room["bookings"].([]interface{})
			if !exists {
				continue // Skip if no bookings
			}

			// Sort bookings by start time
			sort.Slice(bookings, func(i, j int) bool {
				startTime1, _ := time.Parse(time.RFC3339, bookings[i].(map[string]interface{})["start"].(string))
				startTime2, _ := time.Parse(time.RFC3339, bookings[j].(map[string]interface{})["start"].(string))
				return startTime1.Before(startTime2) // Ascending order
			})

			// Save sorted bookings
			room["bookings"] = bookings
		}
	}

	return result
}

func GetNearbyBuildings(lat float32, lng float32, distance float32, data map[string]interface{}) []map[string]interface{} {
	// Extract the buildings from the data
	buildings, exists := data["data"].(map[string]interface{})["buildings"].([]interface{})
	if !exists {
		return nil
	}

	var nearbyBuildings []map[string]interface{}

	for _, building := range buildings {
		buildingMap, ok := building.(map[string]interface{})
		if !ok {
			continue 
		}

		latRaw, latExists := buildingMap["lat"].(float64)
		lngRaw, lngExists := buildingMap["long"].(float64)

		if !latExists || !lngExists {
			continue
		}
		latField := float32(latRaw)
		lngField := float32(lngRaw)
		if (lngField-lng)*(lngField-lng)+(latField-lat)*(latField-lat) > distance*distance {
			continue
		}
		nearbyBuildings = append(nearbyBuildings, buildingMap)
	}

	return nearbyBuildings
}

func GetFreeBuildings(buildings []map[string]interface{}) []map[string]interface{} {
	var freeBuildings []map[string]interface{}

	for _, building := range buildings {
		rooms, exists := building["rooms"].([]interface{})
		if !exists || len(rooms) == 0 {
			continue
		}

		totalRooms := len(rooms)
		freeRooms := 0

		for _, room := range rooms {
			roomMap, ok := room.(map[string]interface{})
			if !ok {
				continue
			}

			bookings, exists := roomMap["bookings"].([]interface{})
			if !exists || len(bookings) == 0 {
				freeRooms++
				continue
			}

			// TO DO: CHANGE LOGIC TO LOOP THROUGH BOOKINGS AND CHECK IF ANY ROOM IS FREE
			lastBooking := bookings[len(bookings)-1].(map[string]interface{})
			endTime, err := time.Parse(time.RFC3339, lastBooking["end"].(string))
			if err != nil {
				continue
			}

			if endTime.Before(time.Now()) {
				freeRooms++
			}
		}

		// Calculate "freeness" score
		freeness := float64(freeRooms) / float64(totalRooms)
		building["freeness"] = roundToSigFigs(freeness * 100, 2)

		// Only add buildings that have at least one free room
		if freeRooms > 0 {
			freeBuildings = append(freeBuildings, building)
		}
	}

	return freeBuildings
}

func roundToSigFigs(val float64, sigFigs int) float64 {
	if val == 0 {
		return 0
	}
	digits := math.Floor(math.Log10(math.Abs(val))) + 1
	roundFactor := math.Pow(10, float64(sigFigs)-digits)
	return math.Round(val*roundFactor) / roundFactor
}

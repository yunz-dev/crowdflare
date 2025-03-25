package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"time"
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

	// Return the result map
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

		latField, latExists := buildingMap["lat"].(float32)
		lngField, lngExists := buildingMap["long"].(float32)

		if !latExists || !lngExists {
			continue
		}
		if (lngField-lng)*(lngField-lng)+(latField-lat)*(latField-lat) > distance*distance {
			continue
		}
		nearbyBuildings = append(nearbyBuildings, buildingMap)
	}

	// Return the filtered buildings
	return nearbyBuildings
}
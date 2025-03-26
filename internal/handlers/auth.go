package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yunz-dev/crowdflare/internal/models"
	"github.com/yunz-dev/crowdflare/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterHandler handles user registration
func RegisterHandler(db *mongo.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var creds models.User
        if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        if creds.Username == "" || creds.Password == "" {
            http.Error(w, "Username and password are required", http.StatusBadRequest)
            return
        }

        err := models.RegisterUser(db, creds.Username, creds.Password)
        if err != nil {
            if err.Error() == "user already exists" {
                http.Error(w, "User already exists", http.StatusConflict)
            } else {
                http.Error(w, "Error registering user", http.StatusInternalServerError)
            }
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
    }
}

// LoginHandler handles user login and JWT issuance
func LoginHandler(db *mongo.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var creds models.User
        if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        if creds.Username == "" || creds.Password == "" {
            http.Error(w, "Username and password are required", http.StatusBadRequest)
            return
        }

        _, err := models.AuthenticateUser(db, creds.Username, creds.Password)
        if err != nil {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        token, err := service.GenerateJWT(creds.Username)
        if err != nil {
            http.Error(w, "Error generating token", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    }
}

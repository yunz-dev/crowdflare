package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yunz-dev/crowdflare/internal/handlers"
	"github.com/yunz-dev/crowdflare/internal/db"
)

// HeartHandler checks if the server is running.
func HeartHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK) // Return status code 200 OK
    fmt.Fprintf(w, "Server is up and running!") // Response message
}


func main() {
    db.ConnectDB()
  	staticDir := "web/static"
	  fs := http.FileServer(http.Dir(staticDir))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/heart", HeartHandler)
    http.HandleFunc("/", handlers.LandingPage)

    // Start the server
    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server: %v\n", err)
    }
  os.Exit(1)
}

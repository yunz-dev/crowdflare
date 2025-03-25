package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yunz-dev/crowdflare/internal/handlers"
	"github.com/yunz-dev/crowdflare/internal/db"
	"github.com/yunz-dev/crowdflare/internal/middleware"
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
    http.Handle("/static/", http.StripPrefix("/static/", fs))  // Serve static files first
    http.HandleFunc("/heart", HeartHandler)


  	http.HandleFunc("/protected", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-User")
		fmt.Fprintf(w, "Welcome, %s! You accessed a protected route.", user)
	}))

    http.HandleFunc("/leaderboard", handlers.LeaderboardPage)
  	http.HandleFunc("/login", handlers.LoginHandler(db.Client.Database("testdb")))
  	http.HandleFunc("/register", handlers.RegisterHandler(db.Client.Database("testdb")))

    http.HandleFunc("/", handlers.LandingPage)
    http.HandleFunc("/leaderboardData", handlers.LeaderboardData)
    http.HandleFunc("/app", handlers.AppPage)
    http.HandleFunc("POST /api/flare", handlers.AddFlare)
    http.HandleFunc("PUT /api/flare/upvote", handlers.UpvoteFlare)
    http.HandleFunc("PUT /api/flare/downvote", handlers.DownvoteFlare)
    http.HandleFunc("GET /api/flares", handlers.GetFlares)

    // Start the server
    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server: %v\n", err)
    }
  os.Exit(1)
}

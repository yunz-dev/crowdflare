package handlers

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yunz-dev/crowdflare/internal/db"
	"github.com/yunz-dev/crowdflare/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func FlarePopup(w http.ResponseWriter, r *http.Request) {
	fid := r.URL.Query().Get("id")
	flares := db.GetCollection("flares")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(fid)
	if err != nil {
		w.WriteHeader(422)
		return
	}
	filter := bson.D{{"_id", id}}
	rawFlare := flares.FindOne(ctx, filter)
	var flare models.Flare
	rawFlare.Decode(&flare)
	data := struct {
		ID        string
		Rating    int
		Upvotes   int
		Downvotes int
		OwnerId   string
	}{fid, flare.Rating, flare.Upvotes, flare.Downvotes, flare.OwnerId}
	tmpl := template.Must(template.ParseFiles("web/templates/partials/flarepopup.html"))
	tmpl.Execute(w, data)
}

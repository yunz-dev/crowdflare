package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yunz-dev/crowdflare/internal/db"
	"github.com/yunz-dev/crowdflare/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func AddFlare(w http.ResponseWriter, r *http.Request) {
    // username, ok := r.Context().Value("username").(string)
    // if !ok {
    //   username = "z123456"
    // }
    username := "z123456"
    err := r.ParseForm()
    if err != nil {
        w.WriteHeader(422)
        return
    }
    rawLat := r.FormValue("lat")
    rawLng := r.FormValue("lng")
    rawRating := r.FormValue("rating")
    lat, err := strconv.ParseFloat(rawLat, 32)
    if err != nil {
        w.WriteHeader(422)
        return
    }
    lng, err := strconv.ParseFloat(rawLng, 32)
    if err != nil {
        w.WriteHeader(422)
        return
    }
    rating, err := strconv.ParseInt(rawRating, 10, 32)
    if err != nil {
        w.WriteHeader(422)
        return
    }
    if lat < -180.0 || lat > 180.0 ||
    lng < -180.0 || lng > 180.0 ||
    rating < 0 || rating > 5 {
        w.WriteHeader(422)
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    flares := db.GetCollection("flares")
    flare := models.Flare{
        Lat:       lat,
        Lng:       lng,
        Rating:    int(rating),
        Upvotes:   0,
        Downvotes: 0,
        OwnerId:   username,
        Fired:     time.Now(),
    }

    _, err = flares.InsertOne(ctx, flare)
    if err != nil {
        w.WriteHeader(500)
        return
    }
}

func UpvoteFlare(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fid := r.FormValue("id")
    flares := db.GetCollection("flares")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id, err := primitive.ObjectIDFromHex(fid)
    if err != nil {
        w.WriteHeader(422)
        return
    }
    filter := bson.D{{"_id", id}}
    update := bson.D{{"$inc", bson.D{{"Upvotes", 1}}}}
    opts := options.Update().SetUpsert(true)
    _, err = flares.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        w.WriteHeader(404)
    }
}

func DownvoteFlare(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fid := r.FormValue("id")
    flares := db.GetCollection("flares")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id, err := primitive.ObjectIDFromHex(fid)
    if err != nil {
        w.WriteHeader(422)
        return
    }
    filter := bson.D{{"_id", id}}
    update := bson.D{{"$inc", bson.D{{"Downvotes", 1}}}}
    opts := options.Update().SetUpsert(true)
    _, err = flares.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        w.WriteHeader(404)
    }
}

func GetFlares(w http.ResponseWriter, r *http.Request) {
    flares := db.GetCollection("flares")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    one_hour := time.Duration(60 * 60 * 1_000_000_000)
    filter := bson.D{{"fired", bson.D{{"$gte", time.Now().Add(-one_hour)}}}}
    flare_cursor, err := flares.Find(ctx, filter)

    var all_flares []models.Flare
    flare_cursor.All(ctx, &all_flares)

    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(all_flares)
}

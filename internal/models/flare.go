package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flare struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
	Lat       float64
	Lng       float64
	Rating    int
	Upvotes   int
	Downvotes int
	OwnerId   string
	Fired     time.Time
}

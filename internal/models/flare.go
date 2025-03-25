package models

import {
	"time"
}

type Flare struct {
	Lat float32
	Lng float32
	Business int
	Upvotes int
	Downvotes int
	OwnerID string
	Fired time.Time
}
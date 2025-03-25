package models

type User struct {
    Username  string `bson:"username"`
    Karma int    `bson:"karma"`
}
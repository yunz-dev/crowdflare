package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    Username      string        `bson:"username"`
    Password      string        `bson:"password"`
    Streak        int           `bson:"streak"`
    FlareCount    int           `bson:"contribution_count"`
    FlareHistory  []Flare       `bson:"contribution_history"`
    Karma         int           `bson:"karma"`
    KarmaHistory  []int         `bson:"karma_history"`
}

// Mock user storage (replace with DB later)
var users = make(map[string]string)

// HashPassword securely hashes a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifies a hashed password
func CheckPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func RegisterUser(db *mongo.Database, username, password string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    users := db.Collection("users")

    // Check if the user already exists
    var existingUser User
    err := users.FindOne(ctx, bson.M{"username": username}).Decode(&existingUser)
    if err == nil {
        return errors.New("user already exists")
    }
    if err != mongo.ErrNoDocuments {
        return err
    }

    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }

    newUser := User{
        Username:      username,
        Password:      hashedPassword,
        Streak:        0,
        FlareCount:    0,
        FlareHistory:  []Flare{},
        Karma:         0,
        KarmaHistory:  []int{},
    }

    _, err = users.InsertOne(ctx, newUser)
    return err
}

func AuthenticateUser(db *mongo.Database, username, password string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    users := db.Collection("users")

    // Find the user in the database
    var user User
    err := users.FindOne(ctx, bson.M{"username": username}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        return false, errors.New("user not found")
    }
    if err != nil {
        return false, err
    }

    // Compare the stored hashed password with the provided password
    err = CheckPassword(user.Password, password)
    return err == nil, err
}

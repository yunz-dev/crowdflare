package models

type LeaderboardData struct {
	TopUsers []User		`bson:top_users`
	Contributors []User	`bson:contributors`
}
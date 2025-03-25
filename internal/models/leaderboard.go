package models

type LeaderboardData struct {
	TopUsers []Users		`bson:top_users`
	Contributors []Users	`bson:contributors`
}
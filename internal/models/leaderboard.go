package models

type LeaderboardData struct {
	CurrentUser User	`bson:current_user`
	TopUsers []User		`bson:top_users`
	Contributors []User	`bson:contributors`
}
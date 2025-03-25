package models

type LeaderboardData struct {
	CurrentUser User	`bson:current_user`
	Users []User		`bson:top_users`
}
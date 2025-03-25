package models

type LeaderboardUser struct {
    Username  string        `bson:"username"`
    Streak int              `bson:"streak"`
    FlareCount int          `bson:"contribution_count"`
    FlareHistory []Flare    `bson:"contribution_history"`
    Karma int               `bson:"karma"`
    KarmaHistory []int      `bson:"karma_history"`
}

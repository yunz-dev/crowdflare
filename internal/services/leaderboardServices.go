package service

import (
	"github.com/yunz-dev/crowdflare/internal/models"
	"sort"
)

func SliceDataByTime(data []models.LeaderboardUser, timeframe string) []models.LeaderboardUser {
	if timeframe == "all" {
		sort.Slice(data, func(i, j int) bool {
			return data[i].Karma > data[j].Karma
		})

		second := data[1]
		data[1] = data[0]
		data[0] = second
		return data
	}
	days := 0
	if timeframe == "month" {
		days = 30
	} else {
		days = 7
	}

	for i := range data {
		if len(data[i].KarmaHistory) < days {
			continue
		}
		latestKarma := data[i].KarmaHistory[len(data[i].KarmaHistory)-1]
		previousKarma := data[i].KarmaHistory[len(data[i].KarmaHistory)-days]
		data[i].Karma = latestKarma - previousKarma
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Karma > data[j].Karma
	})

	second := data[1]
	data[1] = data[0]
	data[0] = second
	return data
}

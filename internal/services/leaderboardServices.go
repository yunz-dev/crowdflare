package services
import (
	"github.com/yunz-dev/crowdflare/internal/models"
	"sort"
)

func SliceDataByTime(data *models.LeaderboardData, timeframe string) {
	if timeframe == "all" {
		return
	}
	var days int
	if timeframe == "month" {
		days = 30
	} else {
		days = 7
	}
		
	for i := range data.Users {
		if len(data.Users[i].KarmaHistory) < days {
			continue
		}
		latestKarma := data.Users[i].KarmaHistory[len(data.Users[i].KarmaHistory)-1]
		previousKarma := data.Users[i].KarmaHistory[len(data.Users[i].KarmaHistory)-days]
		data.Users[i].Karma = latestKarma - previousKarma
	}
	
	sort.Slice(data.Users, func(i, j int) bool {
		return data.Users[i].Karma > data.Users[j].Karma
	})

	second := data.Users[1]
	data.Users[1] = data.Users[0]
	data.Users[0] = second
}
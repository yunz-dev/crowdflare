package services
import (
	"github.com/yunz-dev/crowdflare/internal/models"
)

func RearrangeTop3(data *models.LeaderboardData) {
	if len(data.TopUsers) != 3{
		return
	}
	second := data.TopUsers[1]
	data.TopUsers[1] = data.TopUsers[0]
	data.TopUsers[0] = second
}
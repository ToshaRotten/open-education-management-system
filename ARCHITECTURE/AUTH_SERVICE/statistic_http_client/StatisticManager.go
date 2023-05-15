package statistic_http_client

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/StatisticManager/models"
)

type StatisticManager struct {
	Last models.Stats
}

func New() *StatisticManager {
	return &StatisticManager{}
}

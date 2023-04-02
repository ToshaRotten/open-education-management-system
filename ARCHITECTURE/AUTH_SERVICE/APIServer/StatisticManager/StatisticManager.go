package StatisticManager

import "github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/StatisticManager/models"

type StatisticManager struct {
	Last models.Stats
}

func New() *StatisticManager {
	return &StatisticManager{}
}

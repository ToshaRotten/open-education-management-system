package StatisticManager

import "github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/StatisticManager/models"

type StatisticManager struct {
	Last models.Stats
}

func New() *StatisticManager {
	return &StatisticManager{}
}

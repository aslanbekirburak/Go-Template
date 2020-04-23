package interfaces

import (
	"common_dashboard_backend/entities"
)

type DashboardUseCases interface {
	CreateComponent(data entities.DashboardPanel, redisKey string) *entities.ErrorType
	GetComponent(redisKey string) (data entities.DashboardPanel, er *entities.ErrorType)
}

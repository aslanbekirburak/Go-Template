package interfaces

import (
	"common_dashboard_backend/entities"
)

type DashboardUseCases interface {
	CreateComponent(data string, redisKey string) *entities.ErrorType
	GetComponent(redisKey string) (data string, er *entities.ErrorType)
}

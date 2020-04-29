package interfaces

import (
	"common_dashboard_backend/entities"
)

type DashboardUseCases interface {
	CreateComponent(data, namespace, redisKey string) *entities.ErrorType
	GetComponent(redisKey string) (data string, er *entities.ErrorType)
	GetNamespace() (data []string, er *entities.ErrorType)
	SetNamespace(ns string) (er *entities.ErrorType)
	GetNameKeys(searchKey string) (data []string, er *entities.ErrorType)
	DeleteComponent(namespace, redisKey string) *entities.ErrorType
}

package interfaces

import (
	"common_dashboard_backend/entities"
)

type RedisStorageGateway interface {
	SetRedisTest(key string, data entities.DashboardPanel) error
	GetRedisTest(key string) (data entities.DashboardPanel, er error)
}

package interactors

import (
	. "common_dashboard_backend/common/projectArch/interactors"
	"common_dashboard_backend/entities"
)

type DashboardInteractor struct{}

func (di *DashboardInteractor) CreateComponent(data string, redisKey string) *entities.ErrorType {
	RedisCommStorage.SetRedisTest(redisKey, data)
	if redisKey == "" {
		return GetError(10010)
	}
	return nil
}

func (di *DashboardInteractor) GetComponent(redisKey string) (data string, er *entities.ErrorType) {
	data, err := RedisCommStorage.GetRedisTest(redisKey)

	if err != nil {
		return data, GetError(10010)
	}
	return data, nil
}

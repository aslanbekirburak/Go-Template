package interactors

import (
	. "common_dashboard_backend/common/projectArch/interactors"
	"common_dashboard_backend/entities"
	"fmt"
)

type DashboardInteractor struct{}

func (di *DashboardInteractor) CreateComponent(data entities.DashboardPanel, redisKey string) *entities.ErrorType {
	RedisCommStorage.SetRedisTest(redisKey, data)
	fmt.Println("AAA", redisKey)
	if redisKey == "" {
		return GetError(10010)
	}
	return nil
}

func (di *DashboardInteractor) GetComponent(redisKey string) (data entities.DashboardPanel, er *entities.ErrorType) {
	data, err := RedisCommStorage.GetRedisTest(redisKey)

	if err != nil {
		return data, GetError(10010)
	}
	fmt.Println("ASLANNNNNNN", data)
	return data, nil
}

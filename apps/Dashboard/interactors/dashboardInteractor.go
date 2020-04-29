package interactors

import (
	. "common_dashboard_backend/common/projectArch/interactors"
	"common_dashboard_backend/entities"
	"fmt"
)

type DashboardInteractor struct{}

func (di *DashboardInteractor) CreateComponent(data, namespace, redisKey string) *entities.ErrorType {
	RedisCommStorage.SetRedisTest(redisKey, data)
	if redisKey == "" {
		return GetError(10010)
	}

	err := RedisCommStorage.SetRedisNamespaces(namespace)
	if err != nil {
		fmt.Println(err)
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

func (di *DashboardInteractor) GetNamespace() (data []string, er *entities.ErrorType) {
	data, err := RedisCommStorage.GetRedisNamespaces()

	if err != nil {
		return data, GetError(10010)
	}
	return data, nil
}

func (di *DashboardInteractor) SetNamespace(ns string) (er *entities.ErrorType) {
	err := RedisCommStorage.SetRedisNamespaces(ns)
	if err != nil {
		return GetError(10010)
	}
	return nil
}

func (di *DashboardInteractor) GetNameKeys(searchKey string) (data []string, er *entities.ErrorType) {
	data, err := RedisCommStorage.GetRedisKeys(searchKey)

	if err != nil {
		return data, GetError(10010)
	}
	return data, nil
}

func (di *DashboardInteractor) DeleteComponent(namespace, redisKey string) *entities.ErrorType {
	RedisCommStorage.DeleteLayoutComponent(namespace, redisKey)
	if redisKey == "" {
		return GetError(10010)
	}

	return nil
}

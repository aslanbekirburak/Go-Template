package redis

import (
	"common_dashboard_backend/entities"
	"encoding/json"
	"fmt"

	. "github.com/go-redis/redis"
)

type RedisStorage struct {
	client *Client
}

var rsComm *RedisStorage

func GetRedisCommStorage() *RedisStorage {
	if rsComm == nil {

		rsComm = &RedisStorage{
			client: clientComm,
		}
	}

	return rsComm
}

func (rs RedisStorage) SetRedisTest(key string, data entities.DashboardPanel) error {

	value, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = rs.client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (rs RedisStorage) GetRedisTest(key string) (data entities.DashboardPanel, err error) {

	var dashboardData entities.DashboardPanel

	res := rs.client.Get(key)
	err = json.Unmarshal([]byte(res.Val()), &dashboardData)
	if err != nil {
		fmt.Println(err)
		return dashboardData, err
	}

	return dashboardData, nil
}

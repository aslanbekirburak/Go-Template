package redis

import (
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

func (rs RedisStorage) SetRedisTest(key string, data string) error {

	err := rs.client.Set(key, data, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (rs RedisStorage) GetRedisTest(key string) (data string, err error) {

	res, err := rs.client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

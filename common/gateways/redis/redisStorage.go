package redis

import (
	"fmt"
	"strings"

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

func (rs RedisStorage) GetRedisNamespaces() (data []string, err error) {

	res, err := rs.client.SMembers("namespaces").Result()
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

func (rs RedisStorage) SetRedisNamespaces(namespace string) error {

	err := rs.client.SAdd("namespaces", namespace).Err()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (rs RedisStorage) GetRedisKeys(searchKey string) (data []string, err error) {

	res, err := rs.client.Keys(searchKey + "_*").Result()
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range res {
		fmt.Println(v)
		fmt.Println(strings.Replace(v, searchKey+"_", "", -1))
		res[i] = strings.Replace(v, searchKey+"_", "", -1)
	}
	return res, nil
}

func (rs RedisStorage) DeleteLayoutComponent(namespace, redisKey string) error {

	err := rs.client.Del(namespace + "_" + redisKey).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, _ := rs.GetRedisKeys(namespace)
	if len(val) <= 0 {
		err := rs.client.SRem("namespaces", namespace).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

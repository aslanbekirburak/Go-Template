package redis

import (
	. "common_dashboard_backend/common/logger"

	"github.com/go-redis/redis"
)

var clientComm *redis.Client

func Init(addressComm string, addressProd string, port string) {

	clientComm = redis.NewClient(&redis.Options{
		Addr:     addressComm + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := clientComm.Ping().Result()
	if err != nil {
		LogError(err)
	} else {
		LogInfo(pong)
	}

}

func Close() {
	clientComm.Close()
}

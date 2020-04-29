package redis

import (
	. "common_dashboard_backend/common/logger"
	"fmt"

	"github.com/go-redis/redis"
)

var clientComm *redis.Client

func Init(host string, port string) {
	redisAddr := host + ":" + port
	fmt.Println(redisAddr)
	clientComm = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
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

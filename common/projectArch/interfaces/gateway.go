package interfaces

type RedisStorageGateway interface {
	SetRedisTest(key string, data string) error
	GetRedisTest(key string) (data string, er error)
}

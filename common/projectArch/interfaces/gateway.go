package interfaces

type RedisStorageGateway interface {
	SetRedisTest(key string, data string) error
	GetRedisTest(key string) (data string, er error)
	SetRedisNamespaces(namespace string) error
	GetRedisNamespaces() (data []string, err error)
	GetRedisKeys(searchKey string) (data []string, err error)
	DeleteLayoutComponent(namespace, redisKey string) error
}

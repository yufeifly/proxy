package redis

type Backend interface {
	RedisGet()
	RedisPut()
}

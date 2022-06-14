package infrastracture

import "github.com/go-redis/redis"

type Redis struct {
	client *redis.Client
}

func NewRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return Redis{
		client: client,
	}
}

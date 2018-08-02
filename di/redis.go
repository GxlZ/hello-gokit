package di

import (
	goredis "github.com/go-redis/redis"
	"log"
)

type redisClient struct {
	*goredis.Client
}

func newRedisClient() *redisClient {
	rc := goredis.NewClient(&goredis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := rc.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}

	return &redisClient{
		rc,
	}
}

package Database

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func InitRedis() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	_, err := Redisdb.Ping().Result()
	if err != nil {
		fmt.Println("Con Redis Fail", err)
	}
}

package config

import (
	"ExchangeApp/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis")
	}
	global.RedisDB = RedisClient
}

package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client = initRedisDB()

func initRedisDB() *redis.Client {
	client := connectRedis()

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("❌ Could not connect to Redis")
	}
	log.Println("🚀 Connected to Redis")

	return client
}

func connectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "",
		DB:       0,
	})

	return client
}

func PingRedis() error {
	ctx := context.Background()

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

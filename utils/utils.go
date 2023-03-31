package utils

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	ServerPort    string
	CacheUrl      string
	CachePassword string
	CacheUsername string
	l             *log.Logger
}

func LoadEnv(l *log.Logger) *EnvironmentConfig {
	if err := godotenv.Load(); err != nil {
		l.Fatalln("Error loading env file")
	}

	return &EnvironmentConfig{
		ServerPort:    os.Getenv("SERVER_PORT"),
		CacheUrl:      os.Getenv("CACHE_URL"),
		CachePassword: os.Getenv("CACHE_PASSWORD"),
		CacheUsername: os.Getenv("CACHE_USERNAME"),
		l:             l,
	}
}

func (env *EnvironmentConfig) ConnectToCache() *redis.Client {
	env.l.Println("Starting connection to cache")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     env.CacheUrl,
		Username: env.CacheUsername,
		Password: env.CachePassword,
		DB:       0,
	})

	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		env.l.Println(err)
	} else {
		env.l.Println("Connected to redis")
	}

	return redisClient
}

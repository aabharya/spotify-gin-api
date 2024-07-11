package db

import (
	"github.com/BinDruid/spotify-gin/settings"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitCache() {

	Redis = redis.NewClient(&redis.Options{
		Addr:     settings.Configs.RedisUrl,
		Password: "",
		DB:       0,
	})
}

package connections

import (
	"github.com/BinDruid/go-practice/settings"
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

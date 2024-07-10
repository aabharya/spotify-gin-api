package connections

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var Redis *redis.Client

func InitCache() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
}

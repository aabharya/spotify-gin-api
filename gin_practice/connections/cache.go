package connections

import (
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitCache() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "",
		DB:       0,
	})
}

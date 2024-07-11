package settings

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

type config struct {
	PostgresUrl string
	RedisUrl    string
	CacheTTl    time.Duration
}

var Configs = &config{getPostgresUrl(), getRedisUrl(), 5 * time.Minute}

func getPostgresUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	pgUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbHost, dbUser, dbPass, dbName, dbPort)
	return pgUrl
}

func getRedisUrl() string {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	rdsUrl := fmt.Sprintf("%s:%s", redisHost, redisPort)
	return rdsUrl
}

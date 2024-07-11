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

var dbHost = os.Getenv("DB_HOST")
var dbUser = os.Getenv("DB_USER")
var dbPass = os.Getenv("DB_PASS")
var dbName = os.Getenv("DB_NAME")
var dbPort = os.Getenv("DB_PORT")
var redisHost = os.Getenv("REDIS_HOST")
var redisPort = os.Getenv("REDIS_PORT")

var pgUrl = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbHost, dbUser, dbPass, dbName, dbPort)
var rdsUrl = fmt.Sprintf("%s:%s", redisHost, redisPort)

var Configs = &config{pgUrl, rdsUrl, 5 * time.Minute}

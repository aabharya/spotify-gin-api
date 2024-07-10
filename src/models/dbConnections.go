package models

import (
	rd "github.com/redis/go-redis/v9"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB
var Redis *rd.Client

func ConnectDatabase() {
	dsn := "host=localhost user=go_practice password=123456 dbname=go_practice port=5314 sslmode=disable TimeZone=Asia/Tehran"
	Postgres, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Postgres.AutoMigrate(&Album{})
	Redis = rd.NewClient(&rd.Options{
		Addr:     "localhost:6378",
		Password: "",
		DB:       0,
	})
}

package db

import (
	"github.com/BinDruid/spotify-gin/models"
	"github.com/BinDruid/spotify-gin/settings"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func InitDB() {
	db, err := gorm.Open(postgres.Open(settings.Configs.PostgresUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Album{})
	Postgres = db
}

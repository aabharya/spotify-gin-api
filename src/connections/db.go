package connections

import (
	"github.com/BinDruid/go-practice/models"
	"github.com/BinDruid/go-practice/settings"
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

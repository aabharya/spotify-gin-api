package connections

import (
	"github.com/BinDruid/go-practice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func InitDB() {
	dsn := "host=localhost user=go_practice password=123456 dbname=go_practice port=5314 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Album{})
	Postgres = db
}

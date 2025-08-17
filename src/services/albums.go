package services

import (
	"encoding/json"
	"github.com/BinDruid/spotify-gin/db"
	"github.com/BinDruid/spotify-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	cachedResult, err := db.Redis.Get(c, "albums").Result()
	if err == redis.Nil {
		db.Postgres.Find(&albums)
		data, _ := json.Marshal(albums)
		db.Redis.Set(c, "albums", data, 5*time.Minute)
	} else {
		_ = json.Unmarshal([]byte(cachedResult), &albums)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	album := models.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	db.Postgres.Create(&album)
	db.Redis.Del(c, "albums")
	c.IndentedJSON(http.StatusCreated, album)
}

func GetAlbumByID(c *gin.Context) {
	var album models.Album
	if err := db.Postgres.Where("ID = ?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": album})
}

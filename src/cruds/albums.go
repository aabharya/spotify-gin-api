package cruds

import (
	"context"
	"encoding/json"
	"github.com/BinDruid/go-practice/connections"
	"github.com/BinDruid/go-practice/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func GetAllAlbums(c *gin.Context) {
	ctx := context.Background()
	var albums []models.Album
	cachedResult, err := connections.Redis.Get(ctx, "albums").Result()
	if err == redis.Nil {
		connections.Postgres.Find(&albums)
		data, _ := json.Marshal(albums)
		connections.Redis.Set(ctx, "albums", data, 5*time.Minute)
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
	connections.Postgres.Create(&album)
	c.IndentedJSON(http.StatusCreated, album)
}

func GetAlbumByID(c *gin.Context) {
	var album models.Album
	if err := connections.Postgres.Where("ID = ?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": album})
}

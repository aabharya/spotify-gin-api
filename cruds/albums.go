package cruds

import (
	"github.com/BinDruid/go-practice/connections"
	"github.com/BinDruid/go-practice/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	connections.Postgres.Find(&albums)
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

package htmx

import (
	"fmt"
	"github.com/BinDruid/spotify-gin/db"
	"github.com/BinDruid/spotify-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RenderAllAlbums(c *gin.Context) {
	var albums []models.Album
	db.Postgres.Find(&albums)
	c.HTML(http.StatusOK, "albums.html", gin.H{
		"Albums": albums,
	})
}

func HxAddAllAlbum(c *gin.Context) {
	time.Sleep(1 * time.Second)
	title := c.PostForm("title")
	artist := c.PostForm("artist")
	priceString := c.PostForm("price")
	price, _ := strconv.Atoi(priceString)
	album := models.Album{Title: title, Artist: artist, Price: uint(price)}
	db.Postgres.Create(&album)
	db.Redis.Del(c, "albums")
	htmxResponse := fmt.Sprintf(`<li class="list-group-item bg-primary text-white">%s - %s</li>`, title, artist)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmxResponse))
}

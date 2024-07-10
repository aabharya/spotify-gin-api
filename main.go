package main

import (
	"github.com/BinDruid/go-practice/connections"
	"github.com/BinDruid/go-practice/cruds"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	connections.InitDB()
	connections.InitCache()
}

func main() {
	router := gin.Default()
	router.Static("/static/", "./public/")
	albumsRouter := router.Group("/albums/")
	albumsRouter.GET("/", cruds.GetAllAlbums)
	albumsRouter.GET("/:id/", cruds.GetAlbumByID)
	albumsRouter.POST("/", cruds.CreateAlbum)
	router.GET("/healthcheck/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	err := router.Run(":4030")
	if err != nil {
		log.Fatal(err)
	}
}

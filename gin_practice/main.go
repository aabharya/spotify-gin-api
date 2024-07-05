package main

import (
	"log"
	"net/http"

	"github.com/BinDruid/go-practice/gin_practice/albums"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static/", "./public")
	albumsRouter := router.Group("/albums")
	albumsRouter.GET("/", albums.GetAll)
	albumsRouter.GET("/:id", albums.GetByID)
	albumsRouter.POST("/", albums.Create)
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	err := router.Run(":4030")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/BinDruid/spotify-gin/db"
	"github.com/BinDruid/spotify-gin/middlewares"
	"github.com/BinDruid/spotify-gin/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	db.InitDB()
	db.InitCache()
}

func main() {
	router := gin.Default()
	router.Use(middlewares.RequestIDMiddleware())
	router.Static("/static/", "../public/")
	api := router.Group("/api")
	v1 := api.Group("/v1")
	routes.AddAlbumsRoute(v1, "/albums/")
	router.GET("/healthcheck/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

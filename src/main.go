package main

import (
	"github.com/BinDruid/spotify-gin/db"
	docs "github.com/BinDruid/spotify-gin/docs"
	"github.com/BinDruid/spotify-gin/middlewares"
	"github.com/BinDruid/spotify-gin/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

func init() {
	db.InitPostgres()
	db.InitRedis()
}

//	@title			Minimal Spotify API
//	@version		1.0
//	@description	Yet Another clone of spotify api built with gin.

// @host		localhost:8000
// @BasePath	/api/v1
func main() {
	router := gin.Default()
	router.Use(middlewares.RequestIDMiddleware())
	router.Static("/static/", "../public/")
	api := router.Group("/api")
	v1 := api.Group("/v1")
	routes.AddAlbumsRoute(v1, "/albums")
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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

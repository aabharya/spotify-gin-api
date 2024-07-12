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
	routes.AddSwaggerRoute(router, "/swagger/*any")
	api := router.Group("/api")
	publicRoutes := api.Group("/v1")
	routes.AddUserRoute(publicRoutes, "/users")
	authOnlyRoutes := publicRoutes.Group("")
	authOnlyRoutes.Use(middlewares.AuthMiddleware())
	routes.AddAlbumsRoute(authOnlyRoutes, "/albums")
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

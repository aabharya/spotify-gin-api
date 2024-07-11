package main

import (
	"github.com/BinDruid/go-practice/connections"
	"github.com/BinDruid/go-practice/middlewares"
	"github.com/BinDruid/go-practice/routes"
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
	router.Use(middlewares.RequestIDMiddleware())
	router.Static("/static/", "../public/")
	api := router.Group("/api")
	v1 := api.Group("/v1")
	routes.AddAlbumsRoute(v1)
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

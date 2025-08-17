package routes

import (
	"github.com/BinDruid/spotify-gin/services"
	"github.com/gin-gonic/gin"
)

func AddAlbumsRoute(rg *gin.RouterGroup, path string) {
	albumsRouter := rg.Group(path)
	albumsRouter.GET("/", services.GetAllAlbums)
	albumsRouter.GET("/:id/", services.GetAlbumByID)
	albumsRouter.POST("/", services.CreateAlbum)
}

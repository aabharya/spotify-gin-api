package routes

import (
	"github.com/BinDruid/spotify-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AddAlbumsRoute(rg *gin.RouterGroup, path string) {
	albumsRouter := rg.Group(path)
	albumsRouter.GET("/", controllers.GetAllAlbums)
	albumsRouter.GET("/:id/", controllers.GetAlbumByID)
	albumsRouter.POST("/", controllers.CreateAlbum)
}

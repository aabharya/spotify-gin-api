package routes

import (
	"github.com/BinDruid/spotify-gin/htmx"
	"github.com/gin-gonic/gin"
)

func AddHtmxRoute(rg *gin.Engine, path string) {
	htmxGroup := rg.Group(path)
	htmxGroup.GET("/albums/", htmx.RenderAllAlbums)
	htmxGroup.POST("/albums/", htmx.HxAddAllAlbum)
}

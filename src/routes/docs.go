package routes

import (
	docs "github.com/BinDruid/spotify-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddSwaggerRoute(r *gin.Engine, path string) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET(path, ginSwagger.WrapHandler(swaggerfiles.Handler))
}

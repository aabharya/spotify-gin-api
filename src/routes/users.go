package routes

import (
	"github.com/BinDruid/spotify-gin/services"
	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup, path string) {
	userRouter := rg.Group(path)
	userRouter.POST("/", services.RegisterUser)
	userRouter.POST("/login/", services.Login)
}

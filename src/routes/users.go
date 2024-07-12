package routes

import (
	"github.com/BinDruid/spotify-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup, path string) {
	userRouter := rg.Group(path)
	userRouter.POST("/", controllers.RegisterUser)
	userRouter.POST("/login/", controllers.Login)
}

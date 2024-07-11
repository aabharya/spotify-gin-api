package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.New()
		ctx.Writer.Header().Set("X-Request-Id", id.String())
		ctx.Next()
	}
}

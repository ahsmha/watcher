package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetEventRoutes(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})
}

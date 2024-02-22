package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("bullshit programmer...")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	r := gin.Default()

	watcherRoutes := r.Group("/wr")
	repoRoute := watcherRoutes.Group("/")
	repoRoute.GET("repo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	r.Run(":8080")
}

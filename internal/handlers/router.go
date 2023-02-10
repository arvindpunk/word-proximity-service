package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(
		cors.New(config),
	)
	v1 := r.Group("/v1")
	v1.GET("/get-target-word", HandleGetTargetWord())

	internal := r.Group("/internal")
	internal.POST("/refresh-word-cache", HandleRefreshWordCache())
	return r
}

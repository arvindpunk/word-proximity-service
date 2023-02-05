package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/get-target-word", GetTargetWord())
	return r
}

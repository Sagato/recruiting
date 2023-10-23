package blog_handler

import "github.com/gin-gonic/gin"

type Handler interface {
	SetupRouter(router *gin.Engine) *gin.Engine
}

package statics

import "github.com/gin-gonic/gin"

func RegisterProtected(router *gin.RouterGroup) {
	router.Static("/assets", "./assets")
}

func RegisterPublic(router *gin.RouterGroup) {
	router.Static("/public", "./public")
}

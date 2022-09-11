package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMe(c *gin.Context) {
	userUuid := fmt.Sprint(c.MustGet("userUuid"))
	c.JSON(http.StatusOK, gin.H{"message": "Hello user " + userUuid})
}

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/me", getMe)
}

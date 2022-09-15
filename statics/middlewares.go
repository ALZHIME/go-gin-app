package statics

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func staticAuth(c *gin.Context) {
	session := sessions.Default(c)
	userUuid := session.Get("uuid")
	fmt.Println(string("\033[32m"), c.Request.Host, c.Request.URL.Path, string("\033[0m"))
	if userUuid == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
	}
	c.Set("userUuid", userUuid)
}

func ProtectedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		staticAuth(c)
	}
}

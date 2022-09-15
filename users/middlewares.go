package users

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func getOrGenerateUuid(c *gin.Context) {
	session := sessions.Default(c)
	userUuid := session.Get("uuid")
	if userUuid == nil {
		userUuid = fmt.Sprint(uuid.New())
		session.Set("uuid", userUuid)
		err := session.Save()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"type": "session save error", "message": fmt.Sprint(err)})
		}
	}
	c.Set("userUuid", userUuid)
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		getOrGenerateUuid(c)
	}
}

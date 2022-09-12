package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"net/http"
)

const (
	maximumIdleConnections = 20
	networkMode            = "tcp"
	redisAddress           = "redis:6379"
	redisPassword          = ""
	encryptKey             = "hello world"
)

var session redis.Store

func Init() (redis.Store, error) {
	store, redisError := redis.NewStore(maximumIdleConnections, networkMode, redisAddress, redisPassword, []byte(encryptKey))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	session = store
	return store, redisError
}

func GetSession() redis.Store {
	return session
}

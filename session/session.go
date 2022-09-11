package session

import (
	"github.com/gin-contrib/sessions/redis"
)

const (
	maximumIdleConnections = 10
	networkMode            = "tcp"
	redisAddress           = "redis:6379"
	redisPassword          = ""
	encryptKey             = "secret"
)

var session redis.Store

func Init() (redis.Store, error) {
	store, redisError := redis.NewStore(maximumIdleConnections, networkMode, redisAddress, redisPassword, []byte(encryptKey))
	//store.Options(sessions.Options{
	//	SameSite: http.SameSiteNoneMode,
	//	Secure:   true,
	//})
	session = store
	return store, redisError
}

func GetSession() redis.Store {
	return session
}

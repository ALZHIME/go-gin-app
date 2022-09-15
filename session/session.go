package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"go-app/variables"
	"net/http"
)

var (
	assertPath = variables.AssertPath()

	maximumIdleConnections = variables.MaximumIdleConnections()
	networkMode            = variables.NetworkMode()
	redisAddress           = variables.RedisAddress()
	redisPassword          = variables.RedisPassword()
	encryptKey             = variables.EncryptKey()

	session redis.Store
)

func Init() (redis.Store, error) {
	store, redisError := redis.NewStore(maximumIdleConnections, networkMode, redisAddress, redisPassword, []byte(encryptKey))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		Path:     assertPath,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	session = store
	return store, redisError
}

func GetSession() redis.Store {
	return session
}

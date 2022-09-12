package variables

var maximumIdleConnections int
var networkMode string
var redisAddress string
var redisPassword string
var encryptKey string

func MaximumIdleConnections() int {
	return maximumIdleConnections
}

func NetworkMode() string {
	return networkMode
}

func RedisAddress() string {
	return redisAddress
}

func RedisPassword() string {
	return redisPassword
}

func EncryptKey() string {
	return encryptKey
}

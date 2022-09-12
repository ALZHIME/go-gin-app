//go:build development

package variables

import "fmt"

func init() {
	fmt.Println("Using {development} configuration.")

	maximumIdleConnections = 20
	networkMode = "tcp"
	redisAddress = "redis:6379"
	redisPassword = ""
	encryptKey = "secret"
}

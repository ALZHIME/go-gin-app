//go:build default

package variables

import "fmt"

func init() {
	fmt.Println("Using {default} configuration.")

	assertPath = "/"

	maximumIdleConnections = 0
	networkMode = ""
	redisAddress = ""
	redisPassword = ""
	encryptKey = ""
}

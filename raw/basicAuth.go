package raw

import (
	"example/goTesting/initializers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
}
func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth || user != os.Getenv("USER") || password !=  os.Getenv("PASSWORD") {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	} else {
		fmt.Println("Authentication Complete!")
	}
	
}

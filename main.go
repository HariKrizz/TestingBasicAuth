package main

import (
	"example/goTesting/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

 func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
 }

 func main() {
		r := gin.Default()
	
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to Building RESTful API",})
		})
		r.Run(":5000")

		// r.GET("/test-coverage", meh())
}

// BasicAuth for test coverage report
// func meh() gin.HandlerFunc {
// 	return gin.BasicAuth(gin.Accounts{os.Getenv("USER"):os.Getenv("PASSWORD")})
	
// }




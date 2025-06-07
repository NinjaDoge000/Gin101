package main

import "github.com/gin-gonic/gin"

// c is a gin context variable
// It represents everything about the HTTP request and response in one place.
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {

	/*
		It creates a new router (gin.Engine).
		It automatically attaches two middleware:
		Logger Middleware: logs every request to the console (method, path, status, latency).
		Recovery Middleware: recovers from any panics and returns a 500 error instead of crashing your app.
	*/

	router := gin.Default()

	router.GET("/ping", pingHandler)

	router.Run(":3000")

}

package main

import (
	"fmt"
	controllers "gin/Controllers"
	services "gin/Services"
	internal "gin/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// c is a gin context variable
// It represents everything about the HTTP request and response in one place.

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func userHandler(c *gin.Context) {
	var user = c.Param("user")
	c.JSON(200, gin.H{
		"message": user,
	})
}

func handleSubmit(c *gin.Context) {

	type user struct {
		Email string `json:"email"  binding:"required"`
		Age   int64  `json:"age" binding:"gte=18,lte=100"`
	}

	var u user

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": u.Email,
		"age":   u.Age,
	})
}

func main() {

	/*
		It creates a new router (gin.Engine).
		It automatically attaches two middleware:
		Logger Middleware: logs every request to the console (method, path, status, latency).
		Recovery Middleware: recovers from any panics and returns a 500 error instead of crashing your app.
		A router is a component that maps incoming HTTP requests to the correct handler function in your code.
	*/

	router := gin.Default()

	db := internal.InitDB()

	if db == nil {
		fmt.Print("Unable to connect to DB")
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)

	// basic route
	router.GET("/ping", pingHandler)

	// route with param
	router.GET("/me/:user", userHandler)

	// post
	router.POST("/submit", handleSubmit)

	notesController := &controllers.NotesController{}
	notesController.InitNotesController(notesService)
	notesController.InitRouter(router)

	notesController.GetNotes()
	notesController.CreateNotes()

	authService := services.InitAuthService(db)
	authController := controllers.InitAuthController(authService)
	authController.InitRouter(router)

	router.Run(":3000")

}

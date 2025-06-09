package controllers

import (
	"fmt"
	services "gin/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controllers Handles HTTP requests and delegates business logic to services.

type NotesController struct {
	notesService *services.NotesService
}

func (n *NotesController) InitNotesController(router *gin.Engine, notesService *services.NotesService) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())

	// register service
	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(c *gin.Context) {

		statusStr := c.Query("status")
		status, err := strconv.ParseBool(statusStr)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		notes, err := n.notesService.GetNotesSerivce(status)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"note": notes,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {

	type notesBody struct {
		Title  string `json:"title"`
		Status bool   `Json:"status"`
	}

	return func(c *gin.Context) {

		var notesBody notesBody
		if err := c.BindJSON(&notesBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		note, err := n.notesService.CreateNotesSerivce(notesBody.Title, notesBody.Status)

		if err != nil {
			fmt.Print("Could not create notes in DB")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})
	}
}

package controllers

import (
	services "gin/Services"
	"net/http"

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
		c.JSON(http.StatusOK, gin.H{
			"note": n.notesService.GetNotesSerivce(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"note": n.notesService.CreateNotesSerivce(),
		})
	}
}

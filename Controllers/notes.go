package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct{}

func (n *NotesController) InitNotesController(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"note": "displaying notes",
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"note": "created notes",
		})
	}
}

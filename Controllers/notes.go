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
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())

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

func (n *NotesController) UpdateNotes() gin.HandlerFunc {

	type notesBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}

	return func(c *gin.Context) {

		var notesBody notesBody
		if err := c.BindJSON(&notesBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		note, err := n.notesService.UpdateNotesService(notesBody.Title, notesBody.Status, notesBody.Id)

		if err != nil {
			fmt.Print("Could not create notes in DB")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		notes, err := n.notesService.DeleteNotesService(id)

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

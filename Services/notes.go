package services

import (
	// internal "gin/internal/database"
	internal "gin/internal/models"

	"fmt"

	"gorm.io/gorm"
)

// logic related to connecting to DB, yada, yada, goes here...
type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type notes struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesSerivce() []notes {
	return []notes{
		{Id: 1, Name: "note1"},
		{Id: 2, Name: "note2"},
		{Id: 3, Name: "note3"},
	}
}

func (n *NotesService) CreateNotesSerivce() notes {

	err := n.db.Create(&internal.Notes{
		Id:     1,
		Title:  "Notes",
		Status: true,
	})

	if err != nil {
		fmt.Print(err)
	}

	return notes{
		Id:   1,
		Name: "note1",
	}
}

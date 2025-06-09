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

// type notes struct {
// 	Id   int
// 	Name string
// }

func (n *NotesService) GetNotesSerivce(status bool) ([]*internal.Notes, error) {

	var notes []*internal.Notes

	if err := n.db.Where("status=?", status).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NotesService) CreateNotesSerivce(title string, status bool) (*internal.Notes, error) {

	notes := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if result := n.db.Create(notes); result.Error != nil {
		fmt.Println("error creating notes")
		return nil, result.Error
	}

	return notes, nil
}

func (n *NotesService) UpdateNotesService(title string, status bool, id int) (*internal.Notes, error) {

	var note *internal.Notes
	if result := n.db.Where("Id=?", id).First(&note); result.Error != nil {
		fmt.Println("error creating notes")
		return nil, result.Error
	}

	note.Title = title
	note.Status = status

	if result := n.db.Save(&note); result.Error != nil {
		fmt.Println("error creating notes")
		return nil, result.Error
	}

	return note, nil
}

func (n *NotesService) DeleteNotesService(id int64) (*internal.Notes, error) {

	var notes *internal.Notes

	if err := n.db.Where("id=?", id).First(&notes).Error; err != nil {
		return nil, err
	}

	if err := n.db.Where("id=?", id).Delete(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

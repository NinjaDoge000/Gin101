package internal

type Notes struct {
	Id     int    `gorm:"primaryKey;autoIncrement"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

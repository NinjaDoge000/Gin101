package internal

type Users struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `json:"title"`
	Password string `json:"status"`
}

func (u *Users) TableName() string {
	return "user"
}

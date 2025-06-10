package services

import (
	"errors"
	internal "gin/internal/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(&internal.Users{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) Login(email string, password string) (*internal.Users, error) {

	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	if password == "" {
		return nil, errors.New("password can't be empty")
	}

	var user internal.Users

	user.Email = email
	user.Password = password

	if err := a.db.Where("email=?", email).Where("password=?", password).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthService) Register(email string, password string) (*internal.Users, error) {

	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	if password == "" {
		return nil, errors.New("password can't be empty")
	}

	var user internal.Users

	user.Email = email
	user.Password = password

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

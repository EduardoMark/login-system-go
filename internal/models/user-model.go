package models

import (
	"fmt"
	"time"

	"github.com/EduardoMark/login-system-go/internal/database"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(username, password string) User {
	return User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Save(user *User) error {
	unique, err := isUsernameUnique(user.Username)
	if err != nil {
		return err
	}
	if !unique {
		return fmt.Errorf("username already exists")
	}

	conn := database.Connection()

	s := conn.Save(user)
	if s.Error != nil {
		return s.Error
	}

	return nil
}

func FindOneUser(username string) (*User, error) {
	var user User
	conn := database.Connection()

	result := conn.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func isUsernameUnique(username string) (bool, error) {
	var user User
	conn := database.Connection()

	if err := conn.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

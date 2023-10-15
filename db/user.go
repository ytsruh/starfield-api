package database

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomClaims struct {
	User                 string `json:"user"`
	Id                   uint   `json:"id"`
	jwt.RegisteredClaims `json:"claims"`
}

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  []byte         `json:"password"`
	Keys      []APIKey       `json:"apikeys"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateUser(input *User) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	tx := db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return User{}, tx.Error
	}
	return user, nil
}

func UpdateUser(user *User) error {
	tx := db.Updates(&user)
	return tx.Error
}

func DeleteUser(id string) error {
	deleteId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Unscoped is for a full delete instead of a soft delete
	//tx := db.Unscoped().Delete(&User{}, id)
	tx := db.Delete(&User{}, deleteId)
	return tx.Error
}

package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasswordReset struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewPasswordReset(input *PasswordReset) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func GetPasswordReset(id string) (PasswordReset, error) {
	var passwordReset PasswordReset
	tx := db.Where("id = ?", id).First(&passwordReset)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return PasswordReset{}, tx.Error
	}
	return passwordReset, nil
}

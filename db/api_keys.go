package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIKey struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"name" gorm:"not null"`
	Key       string         `json:"key" gorm:"not null"`
	UserId    uuid.UUID      `json:"owner"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateKey(input *APIKey) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func GetKeysByUser(id string) ([]APIKey, error) {
	ownerId, err := uuid.Parse(id)
	if err != nil {
		return []APIKey{}, err
	}
	var keys []APIKey
	tx := db.Where("user_id = ?", ownerId).Find(&keys)
	if tx.Error != nil {
		return []APIKey{}, tx.Error
	}
	return keys, nil
}

func GetAllKeys() ([]string, error) {
	var keys []string
	result := db.Table("api_keys").Select("Key").Find(&keys)
	if result.Error != nil {
		return nil, result.Error
	}
	return keys, nil
}

func UpdateKey(key *APIKey) error {
	tx := db.Updates(&key)
	return tx.Error
}

func DeleteKey(id string) error {
	deleteId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Unscoped is for a full delete instead of a soft delete
	//tx := db.Unscoped().Delete(&User{}, id)
	tx := db.Delete(&APIKey{}, deleteId)
	return tx.Error
}

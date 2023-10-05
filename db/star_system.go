package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StarSystem struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Read all
func GetAllStarSystems() ([]StarSystem, error) {
	var starSystems []StarSystem
	tx := db.Find(&starSystems)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return []StarSystem{}, tx.Error
	}
	return starSystems, nil
}

// Read one by ID
func GetStarSystem(id string) (StarSystem, error) {
	var starSystem StarSystem
	tx := db.Where("id = ?", id).First(&starSystem)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return StarSystem{}, tx.Error
	}
	return starSystem, nil
}

// Create
func CreateStarSystem(input *StarSystem) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

// Update
func UpdateStarSystem(starSystem *StarSystem) error {
	tx := db.Save(&starSystem)
	return tx.Error
}

// Delete
func DeleteStarSystem(id string) error {
	deleteId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Unscoped is for a full delete instead of a soft delete
	//tx := db.Unscoped().Delete(&StarSystem{}, id)
	tx := db.Delete(&StarSystem{}, deleteId)
	return tx.Error
}

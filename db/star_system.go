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
	Planets   []Planet       `json:"planets"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Read all
func GetAllStarSystems() ([]StarSystem, error) {
	var starSystems []StarSystem
	tx := db.Preload("Planets").Find(&starSystems)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return []StarSystem{}, tx.Error
	}
	return starSystems, nil
}

// Read one by ID
func GetStarSystem(id string) (StarSystem, error) {
	var starSystem StarSystem
	tx := db.Preload("Planets").Where("id = ?", id).First(&starSystem)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return StarSystem{}, tx.Error
	}
	return starSystem, nil
}

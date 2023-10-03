package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StarSystem struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"url" gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetAllStarSystems() ([]StarSystem, error) {
	var starSystems []StarSystem
	tx := db.Find(&starSystems)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return []StarSystem{}, tx.Error
	}
	return starSystems, nil
}

func GetStarSystem(id string) (StarSystem, error) {
	var starSystem StarSystem
	tx := db.Where("id = ?", id).First(&starSystem)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return StarSystem{}, tx.Error
	}
	return starSystem, nil
}

func CreateStarSystem(input *StarSystem) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

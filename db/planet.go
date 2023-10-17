package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Planet struct {
	ID           uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name         string         `json:"name" gorm:"unique;not null"`
	Description  string         `json:"description"`
	StarSystemId uuid.UUID      `json:"starsystem"`
	Moons        []Moon         `json:"moons"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// Read all
func GetAllPlanets() ([]Planet, error) {
	var planets []Planet
	tx := db.Find(&planets)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return []Planet{}, tx.Error
	}
	return planets, nil
}

// Read one by ID
func GetPlanet(id string) (Planet, error) {
	var planet Planet
	tx := db.Where("id = ?", id).First(&planet)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return Planet{}, tx.Error
	}
	return planet, nil
}

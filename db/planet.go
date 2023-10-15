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

// Create
func CreatePlanet(input *Planet) error {
	tx := db.Create(input)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return tx.Error
	}
	return nil
}

// Update
func UpdatePlanet(planet *Planet) error {
	tx := db.Save(&planet)
	return tx.Error
}

// Delete
func DeletePlanet(id string) error {
	deleteId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Unscoped is for a full delete instead of a soft delete
	//tx := db.Unscoped().Delete(&Planet{}, id)
	tx := db.Delete(&Planet{}, deleteId)
	return tx.Error
}

package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Moon struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	PlanetId    uuid.UUID      `json:"planet"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Read all
func GetAllMoons() ([]Moon, error) {
	var moons []Moon
	tx := db.Find(&moons)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return []Moon{}, tx.Error
	}
	return moons, nil
}

// Read one by ID
func GetMoon(id string) (Moon, error) {
	var moon Moon
	tx := db.Where("id = ?", id).First(&moon)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return Moon{}, tx.Error
	}
	return moon, nil
}

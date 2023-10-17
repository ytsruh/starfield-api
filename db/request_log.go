package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestLog struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Path      string         `json:"path" gorm:"not null"`
	Key       string         `json:"key" gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (log *RequestLog) Create() error {
	tx := db.Create(log)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

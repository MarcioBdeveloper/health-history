package models

import "time"

type Symptom struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

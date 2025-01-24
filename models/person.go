package models

import "time"

type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Age       int       `json:"age" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

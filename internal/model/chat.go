package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	*gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	CreatedAt time.Time
	Messages  []Message `gorm:"constraint:OnDelete:CASCADE"`
}

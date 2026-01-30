package model

import (
	"time"
)

type Chat struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	CreatedAt time.Time
	Messages  []Message `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;"` // no extra magic:)
}

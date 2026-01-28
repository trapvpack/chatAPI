package model

import "time"

type Message struct {
	Id        uint   `gorm:"primary_key"`
	ChatId    uint   `gorm:"index"`
	Text      string `gorm:"type:text not null"`
	CreatedAt time.Time
}

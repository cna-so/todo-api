package db

import (
	"time"
)

type Todo struct {
	BaseModel
	Title       string `gorm:"required"`
	Description string
	Status      uint8     `gorm:"required;default:1"`
	ExpireDate  time.Time `json:"expire_date"`
	UserID      string
	TagID       string
}

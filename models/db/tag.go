package db

import (
	"github.com/google/uuid"
)

type Tags struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	UserID      uuid.UUID
	Todos       []Todo `gorm:"foreignKey:TagID"`
}

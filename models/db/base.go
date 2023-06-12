package db

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	CreateAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;not null"`
	UpdateAt time.Time `json:"updated_at,omitempty" gorm:"autoCreateTime;not null"`
}

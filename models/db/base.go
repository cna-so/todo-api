package db

import (
	"time"
)

type BaseModel struct {
	ID       string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	CreateAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;not null"`
	UpdateAt time.Time `json:"updated_at,omitempty" gorm:"autoCreateTime;not null"`
}

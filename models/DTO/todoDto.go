package dto

import (
	"github.com/cna-so/todo-api/models/db"
	"time"
)

type TodoCreateRequestDto struct {
	Title       string                 `json:"title" binding:"required" gorm:"required"`
	Description string                 `json:"description,omitempty"`
	Status      uint8                  `json:"status,omitempty" gorm:"required;default:1"`
	ExpireDate  time.Time              `json:"expire_date,omitempty"`
	Tags        []TodoCreateTagRequest `json:"tags,omitempty"`
	UserID      string                 `json:"user_id,omitempty"`
}
type TodoResponseDto struct {
	ID          string    `json:"id"`
	CreateAt    time.Time `json:"created_at,omitempty" `
	Title       string    `json:"title" `
	Description string    `json:"description"`
	Status      uint8     `json:"status" `
	ExpireDate  time.Time `json:"expire_date"`
	Tags        []*db.Tag `json:"tags"`
}

type TodoCreateTagRequest struct {
	ID string `json:"id"`
}

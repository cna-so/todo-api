package dto

import "time"

type TodoCreateRequestDto struct {
	Title       string    `json:"title" binding:"required" gorm:"required"`
	Description string    `json:"description"`
	Status      uint8     `json:"status" gorm:"required;default:1"`
	ExpireDate  time.Time `json:"expire_date"`
	TagID       string    `json:"tag_id"`
}
type TodoResponseDto struct {
	ID          string    `json:"id"`
	CreateAt    time.Time `json:"created_at,omitempty" `
	Title       string    `json:"title" `
	Description string    `json:"description"`
	Status      uint8     `json:"status" `
	ExpireDate  time.Time `json:"expire_date"`
	TagsID      []string  `json:"tags_id"`
}

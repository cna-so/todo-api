package dto

type TagCreateRequestDto struct {
	Title       string `json:"title" binding:"required"`
	Color       string `json:"color" gorm:"default:black"`
	Description string `json:"description" gorm:"default:null"`
	UserID      string `gorm:"not null"`
}
type TagCreateResponseDto struct {
	ID          string
	Title       string `json:"title"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

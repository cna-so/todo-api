package db

type Tags struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Color       string `json:"color"`
	UserID      string
}

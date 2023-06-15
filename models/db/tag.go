package db

type Tag struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Color       string `json:"color"`
	UserID      string
	Todos       []*Tag `gorm:"many2many:todo_tags"`
}

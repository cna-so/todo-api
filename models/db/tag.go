package db

type Tag struct {
	BaseModel
	Title       string `json:"title" binding:"required" gorm:"not null"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	UserID      string `json:"-"`
	Todos       []*Tag `gorm:"many2many:todo_tags" json:"-"`
}

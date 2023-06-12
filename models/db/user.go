package db

type Users struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"default:null"`
	LastName  string `json:"last_name" gorm:"default:null"`
	Email     string `json:"email" gorm:"unique;not null" binding:"required" `
	Password  string `json:"password" binding:"required" gorm:"unique;not null"`
	Role      string `json:"role"`
	Tags      []Tags `gorm:"foreignKey:UserID"`
	Todos     []Todo `gorm:"foreignKey:UserID"`
}

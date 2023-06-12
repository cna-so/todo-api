package db

type Users struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"default:null"`
	LastName  string `json:"last_name" gorm:"default:null"`
	Email     string `json:"email" binding:"required" gorm:"unique;not null"`
	Password  string `json:"password" binding:"required" gorm:"unique;not null"`
	Role      string `json:"role"`
}

package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/cna-so/todo-api/initializers"
)

type UserModels struct {
	UserUID   string    `json:"user_uid,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdateAt  time.Time `json:"updated_at,omitempty"`
}

func (u *UserModels) CreateUserAccount() (status int, err error) {
	_, err = initializers.Db.Exec(`INSERT INTO users
	(user_uid ,first_name ,last_name , email ,password, role)
	VALUES ($1,$2,$3,$4,$5,$6)`,
		uuid.New(), u.FirstName, u.LastName, u.Email, u.Password, "U") // role is defualt to user we have S as superuser and A as admin
	if err != nil {
		return http.StatusForbidden, err
	}
	return http.StatusOK, nil
}

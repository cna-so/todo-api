package models

import (
	"net/http"
	"time"

	"github.com/cna-so/todo-api/initializers"
)

type UserModels struct {
	UserUID   string    `json:"user_uid,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role,omitempty"`
	CreateAt  time.Time `json:"create_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
}

func (u *UserModels) CreateUserAccount() (status int, err error) {
	_, err = initializers.Db.Exec(`INSERT INTO users
	(user_uid ,first_name ,last_name , email ,password, role , create_at , update_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7)`, u.UserUID, u.FirstName, u.LastName, u.Email, u.Password, u.Role, u.CreateAt, u.UpdateAt)
	if err != nil {
		return http.StatusForbidden, err
	}
	return http.StatusOK, nil
}

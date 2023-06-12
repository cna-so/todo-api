package repository

import (
	"fmt"
	"github.com/cna-so/todo-api/models/db"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// UserRepositoryProvider init user repository
func UserRepositoryProvider(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

// CreateUser create user in sql
func (r *UserRepository) CreateUser(user db.Users) (*db.Users, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*db.Users, error) {
	var user db.Users
	if err := r.db.Model(&db.Users{}).Where(db.Users{Email: email}).First(&user).Error; err != nil {
		fmt.Println(user)
		return nil, err
	} else {
		return &user, nil
	}
}

// CheckUserByEmail check user if exist or not
func (r *UserRepository) CheckUserByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&db.Users{}).Where(db.Users{Email: email}).Count(&count).Error; err != nil {
		return true, err
	}
	return count > 0, nil
}

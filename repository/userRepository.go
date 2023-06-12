package repository

import (
	"github.com/cna-so/todo-api/models/db"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryProvider(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user db.Users) (*db.Users, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUserByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&db.Users{}).Where(db.Users{Email: email}).Count(&count).Error; err != nil {
		return true, err
	}
	return count > 0, nil
}

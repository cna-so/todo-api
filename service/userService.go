package service

import (
	"errors"
	"github.com/cna-so/todo-api/models/db"
	"github.com/cna-so/todo-api/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func UserServiceProvider(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

func (s *UserService) SignUpUser(user db.Users) (*db.Users, error) {
	if isExistUser, err := s.repository.CheckUserByEmail(user.Email); err != nil {
		return nil, err
	} else {
		if isExistUser {
			return nil, errors.New("user already with this email exit")
		}
	}
	if m, err := s.repository.CreateUser(user); err != nil {
		return nil, err
	} else {
		return m, err
	}
}

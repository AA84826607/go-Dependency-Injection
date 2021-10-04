package user

import (
	"go-Dependency-Injection/app/domain/dtos"
	"go-Dependency-Injection/app/domain/dtos/inputs"
	"go-Dependency-Injection/app/domain/repositories"
)

type UserService struct {
	UserRepository repositories.IUserRepository `inject:"UserRepository"`
}

func (s *UserService) AddUser(user *inputs.UserInput) error {
	return s.UserRepository.AddUser(user)
}

func (s *UserService) GetUsers() []dtos.UserDto {
	return s.UserRepository.GetUsers()
}

func (s *UserService) GetUser(id int64) *dtos.UserDto {
	return s.UserRepository.GetUser(id)
}

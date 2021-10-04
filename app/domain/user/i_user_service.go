package user

import (
	"go-Dependency-Injection/app/domain/dtos"
	"go-Dependency-Injection/app/domain/dtos/inputs"
)

type IUserService interface {
	GetUsers() []dtos.UserDto
	GetUser(id int64) *dtos.UserDto
	AddUser(user *inputs.UserInput) error
}

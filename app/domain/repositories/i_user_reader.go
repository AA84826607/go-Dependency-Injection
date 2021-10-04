package repositories

import "go-Dependency-Injection/app/domain/dtos"

type IUserReader interface {
	GetUsers() []dtos.UserDto
	GetUser(id int64) *dtos.UserDto
	GetMaxUserId() int64
}

package repositories

import (
	"go-Dependency-Injection/app/domain/dtos/inputs"
)

type IUserRepository interface {
	IUserReader
	AddUser(user *inputs.UserInput) error
	UpdateUserNickName(id int64, nickName string) error
}

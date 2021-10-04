package user

import (
	"errors"
	"go-Dependency-Injection/app/dal/db"
	"go-Dependency-Injection/app/dal/entities"
	"go-Dependency-Injection/app/domain/dtos/inputs"
)

type UserRepository struct {
	UserRead
	WriteDb *db.MockDB `inject:"MockDBWrite"`
}

func (w *UserRepository) AddUser(user *inputs.UserInput) error {
	model := entities.UserEntity{}
	model.ID = w.GetMaxUserId() + 1
	model.Name = user.Name
	model.NickName = user.NickName
	model.Gender = user.Gender
	model.Age = user.Age
	model.Address = user.Address
	if w.ReadDb.Connect() {
		users := w.ReadDb.Users()
		users = append(users, model)
	}
	return nil
}

func (w *UserRepository) UpdateUserNickName(id int64, nickName string) error {
	user := w.GetUser(id)
	if user.ID > 0 {
		user.NickName = nickName
		return nil
	} else {
		return errors.New("未找到用户信息")
	}
}

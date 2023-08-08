package controller

import (
	"github.com/mb-cdev/p1-auth-service/usecase/user/inputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/outputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/usecase"
)

type User struct {
	usecase usecase.UserInterface
}

func NewUser(usecase usecase.UserInterface) *User {
	return &User{
		usecase: usecase,
	}
}

func (u *User) RegisterAction(data RequestDataInterface) *outputdata.CreateUserOutputData {
	out := u.usecase.CreateUser(
		&inputdata.CreateUserInputData{
			Name:     data.Get("name"),
			Password: data.Get("password"),
		},
	)

	return out
}

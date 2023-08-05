package controller

import "github.com/mb-cdev/p1-auth-service/usecase/user/usecase"

type User struct {
	usecase usecase.UserInterface
}

func NewUser(usecase usecase.UserInterface) *User {
	return &User{
		usecase: usecase,
	}
}

func (u *User) RegisterAction() {

}

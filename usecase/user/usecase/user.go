package usecase

import (
	"github.com/mb-cdev/p1-auth-service/entity"
	"github.com/mb-cdev/p1-auth-service/usecase/user/inputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/repository"
)

type UserInterface interface {
	CreateUser(*inputdata.CreateUserInputData)
	LoginUser()
}

type User struct {
	db repository.UserRepository
}

func NewUser(db repository.UserRepository) UserInterface {
	return &User{
		db: db,
	}
}

func (u *User) CreateUser(input *inputdata.CreateUserInputData) {
	if u.db.IsNameUsed(input.Name) {
		return
	}

	usr := entity.NewUser(input.Name, input.Password)
	u.db.Save(usr)
}

func (u *User) LoginUser() {

}

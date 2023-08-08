package usecase

import (
	"errors"
	"net/http"

	"github.com/mb-cdev/p1-auth-service/entity"
	"github.com/mb-cdev/p1-auth-service/usecase/user/inputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/outputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/repository"
)

var errNameInUse = errors.New("NAME_IN_USE")

type UserInterface interface {
	CreateUser(*inputdata.CreateUserInputData) *outputdata.CreateUserOutputData
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

func (u *User) CreateUser(input *inputdata.CreateUserInputData) *outputdata.CreateUserOutputData {
	if u.db.IsNameUsed(input.Name) {
		return &outputdata.CreateUserOutputData{
			Err:            errNameInUse,
			IsSuccess:      false,
			HttpStatusCode: http.StatusBadRequest,
		}
	}

	usr := entity.NewUser(input.Name, input.Password)
	saved := u.db.Save(usr)

	httpStatusCode := http.StatusCreated
	if !saved {
		httpStatusCode = http.StatusBadRequest
	}

	return &outputdata.CreateUserOutputData{
		IsSuccess:      saved,
		HttpStatusCode: httpStatusCode,
		CreatedId:      usr.ID,
	}
}

func (u *User) LoginUser() {

}

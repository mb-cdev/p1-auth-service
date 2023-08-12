package usecase

import (
	"errors"
	"net/http"

	"github.com/mb-cdev/jwt"
	"github.com/mb-cdev/p1-auth-service/entity"
	"github.com/mb-cdev/p1-auth-service/settings"
	"github.com/mb-cdev/p1-auth-service/usecase/user/inputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/outputdata"
	"github.com/mb-cdev/p1-auth-service/usecase/user/repository"
)

var errNameInUse = errors.New("NAME_IN_USE")
var errBadCredentials = errors.New("ERR_BAD_CREDENTIALS")

type UserInterface interface {
	CreateUser(*inputdata.CreateUserInputData) *outputdata.CreateUserOutputData
	LoginUser(*inputdata.LoginUserInputData) *outputdata.LoginUserOutputData
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

func (u *User) LoginUser(input *inputdata.LoginUserInputData) *outputdata.LoginUserOutputData {
	usr := u.db.FindByName(input.Name)
	if usr == nil {
		return &outputdata.LoginUserOutputData{
			Err:            errBadCredentials,
			HttpStatusCode: http.StatusUnauthorized,
		}
	}

	if !usr.IsPasswordValid(input.Password) {
		return &outputdata.LoginUserOutputData{
			Err:            errBadCredentials,
			HttpStatusCode: http.StatusUnauthorized,
		}
	}

	tkn := jwt.New(jwt.AlgoHS512, settings.JwtSecret)
	tkn.SetPayload("UUID", usr.ID)
	tkn.SetPayload("login", usr.Name)

	return &outputdata.LoginUserOutputData{
		HttpStatusCode: http.StatusAccepted,
		IsSuccess:      true,
		JwtToken:       tkn.String(),
	}
}

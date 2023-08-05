package usecase

import (
	"github.com/mb-cdev/authenticationService/entity"
	"github.com/mb-cdev/authenticationService/usecase/user/inputdata"
	"github.com/mb-cdev/authenticationService/usecase/user/repository"
)

func CreateUser(data inputdata.CreateUserInputDataInterface, userRepository repository.UserRepository) {
	if userRepository.IsNameUsed(data.Name()) {
		return
	}

	user := entity.NewUser(data.Name(), data.Password())
	userRepository.Save(user)
}

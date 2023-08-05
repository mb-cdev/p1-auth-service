package usecase_test

import (
	"testing"

	"github.com/mb-cdev/authenticationService/gateway"
	"github.com/mb-cdev/authenticationService/usecase/user/inputdata"
	"github.com/mb-cdev/authenticationService/usecase/user/usecase"
)

func TestCreateUser(t *testing.T) {
	db := gateway.NewUserMemoryDb()

	testingTable := []struct {
		name              string
		password          string
		expectedIsCreated bool
	}{
		{
			name:              "T1",
			password:          "PASSWD1",
			expectedIsCreated: true,
		},
		{
			name:              "T1",
			password:          "PASSWD1",
			expectedIsCreated: false,
		},
	}

	for _, test := range testingTable {
		input := &inputdata.CreateUserInputData{}
		input.SetName(test.name)
		input.SetPassword(test.password)

		usecase.CreateUser(input, db)
	}

}

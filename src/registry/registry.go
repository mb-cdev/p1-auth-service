package registry

import (
	"github.com/mb-cdev/p1-auth-service/controller"
	"github.com/mb-cdev/p1-auth-service/gateway"
	"github.com/mb-cdev/p1-auth-service/usecase/user/usecase"
)

var ControllerRegistry controllerRegistry

type controllerRegistry struct {
	User *controller.User
}

func (cr *controllerRegistry) Init() {
	cr.User = controller.NewUser(
		usecase.NewUser(
			gateway.NewUserMemoryDb(),
		),
	)
}

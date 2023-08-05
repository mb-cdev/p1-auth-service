package repository

import "github.com/mb-cdev/authenticationService/entity"

type UserRepository interface {
	IsNameUsed(name string) bool
	FindByName(name string) *entity.User
	Save(user *entity.User) bool
	GetAll() []*entity.User
}

package repository

import "github.com/mb-cdev/p1-auth-service/entity"

type UserRepository interface {
	IsNameUsed(name string) bool
	FindByName(name string) *entity.User
	Save(user *entity.User) bool
	GetAll() []*entity.User
}

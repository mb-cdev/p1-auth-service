package entity

import (
	"crypto/sha512"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Password  []byte `json:"-"`
	CreatedAt time.Time
}

func NewUser(name string, password string) *User {
	user := &User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	user.SetPassword(password)
	return user
}

func (u *User) SetPassword(password string) {
	u.Password = hashUserPassword(password)
}

func hashUserPassword(password string) []byte {
	fn := sha512.New()
	fn.Write([]byte(password))
	return fn.Sum(nil)
}

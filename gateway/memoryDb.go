package gateway

import (
	"strings"
	"sync"

	"github.com/mb-cdev/authenticationService/entity"
)

type UserMemoryDb struct {
	mutex sync.RWMutex
	users map[string]*entity.User
}

func NewUserMemoryDb() *UserMemoryDb {
	return &UserMemoryDb{
		users: make(map[string]*entity.User),
	}
}

func (m *UserMemoryDb) IsNameUsed(name string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, u := range m.users {
		if strings.Compare(u.Name, name) == 0 {
			return true
		}
	}

	return false
}

func (m *UserMemoryDb) FindByName(name string) *entity.User {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, u := range m.users {
		if strings.Compare(u.Name, name) == 0 {
			return u
		}
	}

	return nil
}

func (m *UserMemoryDb) Save(user *entity.User) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.users[user.ID]; !ok {
		m.users[user.ID] = user
		return true
	}

	return false
}

func (m *UserMemoryDb) GetAll() []*entity.User {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	us := make([]*entity.User, len(m.users))

	for _, u := range m.users {
		us = append(us, u)
	}

	return us
}

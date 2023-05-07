package memory

import (
	"errors"
	"sync"

	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"
	"github.com/google/uuid"
)

type memoryRepository struct {
	users map[uuid.UUID]domain.User
	sync.Mutex
}

func NewUserRepository() ports.UserRepository {
	return &memoryRepository{
		users: make(map[uuid.UUID]domain.User, 0),
	}
}

func (mr memoryRepository) Get(userID uuid.UUID) (domain.User, error) {
	var user domain.User
	if _, ok := mr.users[userID]; !ok {
		return user, errors.New("user not found")
	}

	user = mr.users[userID]

	return user, nil
}

func (mr memoryRepository) GetAll() []domain.User {

	users := make([]domain.User, 0)

	for _, user := range mr.users {
		users = append(users, user)
	}

	return users
}

func (mr memoryRepository) Update(user domain.User) error {
	return nil
}

func (mr memoryRepository) Create(user domain.User) error {

	if mr.users == nil {

		mr.users = make(map[uuid.UUID]domain.User, 0)

	}

	if _, ok := mr.users[user.ID]; ok {
		return errors.New("user exist")
	}

	mr.Mutex.Lock()
	mr.users[user.ID] = user
	mr.Mutex.Unlock()

	return nil
}

func (mr memoryRepository) FindDuplicate(userName, email string) (exist bool) {

	return false

}

func (mr memoryRepository) Login(userName string) (user domain.User, err error) {

	return

}

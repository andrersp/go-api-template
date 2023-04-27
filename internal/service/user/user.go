package service

import (
	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/domain/user"
	"github.com/andrersp/go-api-template/internal/repository/database"

	"github.com/google/uuid"
)

type ServiceUserConfiguration func(us *serviceUser) error

type ServiceUser interface {
	CreateUser(user.User) (uuid.UUID, error)
	GetUser(uuid.UUID) (user.User, error)
	GetUsers() []user.User
}

type serviceUser struct {
	userRepo user.UserInterface
}

func NewUserService(cfgs ...ServiceUserConfiguration) (ServiceUser, error) {

	us := &serviceUser{}

	for _, cfg := range cfgs {
		err := cfg(us)
		if err != nil {
			return nil, err
		}
	}

	return us, nil
}

func ServiceWithRDB() ServiceUserConfiguration {
	return func(us *serviceUser) error {

		conn, err := config.ConnectDB()

		if err != nil {
			return err
		}

		rdbUser := database.NewDBUser(conn)
		us.userRepo = rdbUser

		return nil

	}
}

func (us serviceUser) CreateUser(user user.User) (userID uuid.UUID, err error) {

	err = user.Validate()
	if err != nil {
		return
	}

	err = us.userRepo.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}
	userID = user.ID
	return
}

func (us serviceUser) GetUser(ID uuid.UUID) (user user.User, err error) {
	return us.userRepo.GetUser(ID)

}

func (us serviceUser) GetUsers() (users []user.User) {
	return us.userRepo.GetUsers()
}

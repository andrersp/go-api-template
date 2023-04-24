package service

import (
	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/domain/user"
	"github.com/andrersp/go-api-template/internal/domain/user/rdb"

	"github.com/google/uuid"
)

type ServiceUserConfiguration func(us *ServiceUser) error

type ServiceUser struct {
	userRepo user.UserRepository
}

func NewUserService(cfgs ...ServiceUserConfiguration) (*ServiceUser, error) {

	us := &ServiceUser{}

	for _, cfg := range cfgs {
		err := cfg(us)
		if err != nil {
			return nil, err
		}
	}

	return us, nil
}

func ServiceWithRDB() ServiceUserConfiguration {
	return func(us *ServiceUser) error {

		conn, err := config.ConnectDB()

		if err != nil {
			return err
		}

		rdbUser := rdb.NewRDBUser(conn)
		us.userRepo = rdbUser

		return nil

	}
}

func (us ServiceUser) CreateUser(user user.User) (userID uuid.UUID, err error) {

	err = us.userRepo.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}
	userID = user.GetId()
	return
}

func (us ServiceUser) GetUser(ID uuid.UUID) (user user.User, err error) {
	return us.userRepo.GetUser(ID)

}

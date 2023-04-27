package service

import (
	"log"
	"testing"

	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/domain/user"
	"github.com/google/uuid"

	"gopkg.in/go-playground/assert.v1"
)

func init_user() {
	err := config.CreateSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = config.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
}
func TestUserService(t *testing.T) {

	init_user()
	user := user.User{
		UserName: "rspandre",
		Email:    "meuemail@mal.com",
		Password: "minhasenha",
	}

	serviceUser, _ := NewUserService(ServiceWithRDB())

	var ID uuid.UUID

	t.Run("TesCreateUser", func(t *testing.T) {
		userID, err := serviceUser.CreateUser(user)

		ID = userID

		assert.Equal(t, err, nil)

	})

	t.Run("TestGetUser", func(t *testing.T) {

		_, err := serviceUser.GetUser(ID)
		assert.Equal(t, err, nil)

	})

	t.Run("TestGetAllUsers", func(t *testing.T) {

		users := serviceUser.GetUsers()

		assert.NotEqual(t, len(users), 0)

	})

}

package service

import (
	"log"
	"testing"

	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/domain/user"

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
	user, _ := user.CreateNewUser("rspandre", "email@mail.com", "minhasenha1234")

	serviceUser, _ := NewUserService(ServiceWithRDB())

	t.Run("TesCreateUser", func(t *testing.T) {
		userID, err := serviceUser.CreateUser(user)

		assert.Equal(t, err, nil)
		assert.Equal(t, user.GetId(), userID)

	})

	t.Run("TestGetUser", func(t *testing.T) {

		_, err := serviceUser.GetUser(user.GetId())
		assert.Equal(t, err, nil)

	})

}

package domain

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestUser(t *testing.T) {
	type testCase struct {
		Name          string
		UserName      string
		Email         string
		Password      string
		ExpectedError error
	}
	testCases := []testCase{
		{
			Name:          "Success",
			UserName:      "myusername",
			Email:         "myemail@mail.com",
			Password:      "MyPassword!@#",
			ExpectedError: nil,
		},
		{
			Name:          "ErrorName",
			UserName:      "",
			Email:         "myemail@mail.com",
			Password:      "MyPassword!@#",
			ExpectedError: ErrUserNameEmpyt,
		},
		{
			Name:          "ErrorEmail",
			UserName:      "myusername",
			Email:         "myemail.com",
			Password:      "MyPassword!@#",
			ExpectedError: ErrInvalidEmail,
		},
		{
			Name:          "ErrorEmailMin",
			UserName:      "myusername",
			Email:         "myemail@mail.com",
			Password:      "12344",
			ExpectedError: ErrInvalidPassword,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			user := User{
				UserName: test.UserName,
				Email:    test.Email,
				Password: test.Password,
			}

			err := user.Validate()

			assert.Equal(t, err, test.ExpectedError)

		})
	}

}

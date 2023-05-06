package domain

import (
	"testing"

	apperrors "github.com/andrersp/go-api-template/pkg/app-errors"
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
			ExpectedError: &apperrors.AppError{Msg: "userName cant be empty"},
		},
		{
			Name:          "ErrorEmail",
			UserName:      "myusername",
			Email:         "myemail.com",
			Password:      "MyPassword!@#",
			ExpectedError: &apperrors.AppError{Msg: "invalid email"},
		},
		{
			Name:          "ErrorEmailMin",
			UserName:      "myusername",
			Email:         "myemail@mail.com",
			Password:      "12344",
			ExpectedError: &apperrors.AppError{Msg: "character number less than 6"},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			_, err := NewUser(
				test.UserName,
				test.Email,
				test.Password,
			)

			assert.Equal(t, err, test.ExpectedError)

		})
	}

}

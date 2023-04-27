package user

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestUser(t *testing.T) {

	type testCase struct {
		TestName string
		UserName string
		Email    string
		Password string
		expected error
	}

	testCases := []testCase{
		{
			TestName: "success",
			UserName: "Andre Luis",
			Email:    "meuemail@mail.com",
			Password: "senha123",
			expected: nil,
		},
		{
			TestName: "errorEmptyUserName",
			UserName: "",
			Email:    "meuemail@mail.com",
			Password: "senha123",
			expected: ErrEmptyName,
		},
		{
			TestName: "errorEmptyEmail",
			UserName: "andre",
			Email:    "",
			Password: "senha123",
			expected: ErrInvalidEmail,
		},
		{
			TestName: "erroInvalidPassword",
			UserName: "andre",
			Email:    "meuemail@mai.com",
			Password: "",
			expected: ErrInvalidPassword,
		},
	}

	for _, test := range testCases {
		t.Run(test.TestName, func(t *testing.T) {

			user := User{
				UserName: test.UserName,
				Email:    test.Email,
				Password: test.Password,
			}
			err := user.Validate()

			assert.Equal(t, test.expected, err)
		})
	}

}

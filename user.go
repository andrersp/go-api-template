package template

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	UserName string
	Email    string
	Password string
}

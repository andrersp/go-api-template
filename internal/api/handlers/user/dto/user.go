package dto

import (
	"github.com/google/uuid"
)

type DtoUserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
}

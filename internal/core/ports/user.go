package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/google/uuid"
)

type UserSerice interface {
	Get(uuid.UUID) (dto.UserResponse, error)
	GetAll() []dto.UserResponse
	Update(dto.UserRequest) error
}

type UserRepository interface {
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(domain.User) error
}

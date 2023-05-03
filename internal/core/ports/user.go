package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/google/uuid"
)

type UserSerice interface {
	Create(dto.UserRequest) error
	Get(uuid.UUID) (dto.UserResponse, error)
	GetAll() []dto.UserResponse
	Update(dto.UserRequest) error
	Login(string, string) (domain.User, error)
}

type UserRepository interface {
	Create(domain.User) error
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(domain.User) error
	FindDuplicate(string, string) bool
	Login(string) (domain.User, error)
}

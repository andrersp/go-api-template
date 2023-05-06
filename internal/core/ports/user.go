package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/google/uuid"
)

type UserSerice interface {
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(username string) error
}

type UserRepository interface {
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(domain.User) error
}

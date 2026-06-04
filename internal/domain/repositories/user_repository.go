package repositories

import (
	"context"
	"real-estate-agency-onion/internal/domain/entities"
)

type UserFilters struct {
	IsActive      *bool
	EmailVerified *bool
	Role          *string
}

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	List(ctx context.Context, filters UserFilters, pagination Pagination) ([]*entities.User, int64, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id int64) error
}

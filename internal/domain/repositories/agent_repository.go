package repositories

import (
	"context"
	"real-estate-agency-onion/internal/domain/entities"
	"real-estate-agency-onion/internal/domain/valueobjects"
)

type Pagination struct {
	Limit  int
	Offset int
}

type AgentFilters struct {
	IsActive *bool
	HasUser  *bool
}

type AgentRepository interface {
	Create(ctx context.Context, agent *entities.Agent) error
	GetByID(ctx context.Context, id int64) (*entities.Agent, error)
	List(ctx context.Context, filters AgentFilters, pagination Pagination) ([]*entities.Agent, int64, error)
	Update(ctx context.Context, agent *entities.Agent) error
	SoftDelete(ctx context.Context, id int64) error
	GetByUserID(ctx context.Context, userID int64) (*entities.Agent, error)
	GetByPhone(ctx context.Context, phone valueobjects.PhoneNumber) (*entities.Agent, error)
}

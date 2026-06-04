package repositories

import (
	"context"
	"real-estate-agency-onion/internal/domain/entities"
	"real-estate-agency-onion/internal/domain/enums"
	"time"
)

type ShowingFilters struct {
	PropertyID *int64
	AgentID    *int64
	Status     *enums.ShowingStatus
	StartFrom  *time.Time
	StartTo    *time.Time
}

type ShowingRepository interface {
	Create(ctx context.Context, showing *entities.Showing) error
	GetByID(ctx context.Context, id int64) (*entities.Showing, error)
	List(ctx context.Context, filters ShowingFilters, pagination Pagination) ([]*entities.Showing, int64, error)
	Update(ctx context.Context, showing *entities.Showing) error
	SoftDelete(ctx context.Context, id int64) error
	GetByAgentID(ctx context.Context, agentID int64, startFrom, endTo time.Time) ([]*entities.Showing, error)
	GetByPropertyID(ctx context.Context, propertyID int64) ([]*entities.Showing, error)
}

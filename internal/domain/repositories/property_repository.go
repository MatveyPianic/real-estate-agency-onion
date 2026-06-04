package repositories

import (
	"context"
	"real-estate-agency-onion/internal/domain/entities"
	"real-estate-agency-onion/internal/domain/enums"
)

type PropertyFilters struct {
	DistrictID   *int64
	AgentID      *int64
	MinPrice     *int64
	MaxPrice     *int64
	MinArea      *int
	MaxArea      *int
	Rooms        *int
	PropertyType *enums.PropertyType
	DealType     *enums.DealType
	Status       *enums.PropertyStatus
}

type PropertyRepository interface {
	// Create создает новый объект недвижимости
	Create(ctx context.Context, property *entities.Property) error

	// get by id получает объект по id
	GetByID(ctx context.Context, id int64) (*entities.Property, error)

	// List получает список объектов с фильтрами и пагинацией
	List(ctx context.Context, filters PropertyFilters, pagination Pagination) ([]*entities.Property, int64, error)

	// Update обновляет объект
	Update(ctx context.Context, property *entities.Property) error

	// SoftDelete мягко удаляет объект
	SoftDelete(ctx context.Context, id int64) error

	// GetByAgentID получает все объекты агента
	GetByAgentID(ctx context.Context, agentID int64) ([]*entities.Property, error)
}

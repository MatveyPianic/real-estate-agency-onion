package repositories

import (
	"context"
	"real-estate-agency-onion/internal/domain/entities"
	"real-estate-agency-onion/internal/domain/enums"
)

type InquiryFilters struct {
	PropertyID *int64
	Status     *enums.InquiryStatus
}

type InquiryRepository interface {
	Create(ctx context.Context, inquiry *entities.Inquiry) error
	GetByID(ctx context.Context, id int64) (*entities.Inquiry, error)
	List(ctx context.Context, filters InquiryFilters, pagination Pagination) ([]*entities.Inquiry, int64, error)
	Update(ctx context.Context, inquiry *entities.Inquiry) error
	SoftDelete(ctx context.Context, id int64) error
	GetByPropertyID(ctx context.Context, propertyID int64) ([]*entities.Inquiry, error)
}

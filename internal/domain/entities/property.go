package entities

import (
	"time"

	"real-estate-agency-onion/internal/domain/enums"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/valueobjects"
)

type Property struct {
	id           int64
	agentID      int64
	districtID   int64
	title        string
	description  string
	price        valueobjects.Money
	area         int
	rooms        int
	propertyType enums.PropertyType
	dealType     enums.DealType
	status       enums.PropertyStatus
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

func NewProperty(
	agentID int64,
	districtID int64,
	title string,
	description string,
	price valueobjects.Money,
	area int,
	rooms int,
	propertyType enums.PropertyType,
	dealType enums.DealType,
) (*Property, error) {
	if agentID <= 0 || districtID <= 0 {
		return nil, domainerrors.ErrInvalidInput
	}

	if title == "" {
		return nil, domainerrors.ErrInvalidInput
	}

	if !propertyType.IsValid() || !dealType.IsValid() {
		return nil, domainerrors.ErrInvalidInput
	}

	return &Property{
		agentID:      agentID,
		districtID:   districtID,
		title:        title,
		description:  description,
		price:        price,
		area:         area,
		rooms:        rooms,
		propertyType: propertyType,
		dealType:     dealType,
		status:       enums.PropertyStatusDraft,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

// геттеры
func (p *Property) ID() int64                        { return p.id }
func (p *Property) AgentID() int64                   { return p.agentID }
func (p *Property) DistrictID() int64                { return p.districtID }
func (p *Property) Title() string                    { return p.title }
func (p *Property) Description() string              { return p.description }
func (p *Property) Price() valueobjects.Money        { return p.price }
func (p *Property) Area() int                        { return p.area }
func (p *Property) Rooms() int                       { return p.rooms }
func (p *Property) PropertyType() enums.PropertyType { return p.propertyType }
func (p *Property) DealType() enums.DealType         { return p.dealType }
func (p *Property) Status() enums.PropertyStatus     { return p.status }
func (p *Property) CreatedAt() time.Time             { return p.createdAt }
func (p *Property) UpdatedAt() time.Time             { return p.updatedAt }
func (p *Property) DeletedAt() *time.Time            { return p.deletedAt }

// бизнес логика
func (p *Property) Publish() error {
	if p.status == enums.PropertyStatusArchived {
		return domainerrors.ErrForbidden
	}
	p.status = enums.PropertyStatusPublished
	p.updatedAt = time.Now()
	return nil
}

func (p *Property) Archive() error {
	p.status = enums.PropertyStatusArchived
	p.updatedAt = time.Now()
	return nil
}

func (p *Property) UpdatePrice(newPrice valueobjects.Money) {
	p.price = newPrice
	p.updatedAt = time.Now()
}

func (p *Property) SoftDelete() {
	now := time.Now()
	p.deletedAt = &now
}

func (p *Property) IsDeleted() bool {
	return p.deletedAt != nil
}

// SetID для repository
func (p *Property) SetID(id int64) { p.id = id }

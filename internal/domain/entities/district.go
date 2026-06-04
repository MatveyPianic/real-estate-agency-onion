package entities

import (
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"strings"
	"time"
)

// структура
type District struct {
	id        int64
	cityID    int64
	name      string
	createdAt time.Time
}

// конструктор
func NewDistrict(name string, cityID int64) (*District, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, domainerrors.ErrInvalidInput // валидация
	}
	return &District{
		name:      name,
		cityID:    cityID,
		createdAt: time.Now(),
	}, nil
}

// геттеры
func (d *District) ID() int64            { return d.id }
func (d *District) CityID() int64        { return d.cityID }
func (d *District) Name() string         { return d.name }
func (d *District) CreatedAt() time.Time { return d.createdAt }

// SetID нужен для repository когда получаем из бд и нужно установить айди
func (d *District) SetID(id int64) { d.id = id }

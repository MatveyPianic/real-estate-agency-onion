package entities

import (
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"strings"
	"time"
)

// структура
type City struct {
	id        int64
	name      string
	createdAt time.Time
}

// конструктор
func NewCity(name string) (*City, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, domainerrors.ErrInvalidInput // валидация
	}
	return &City{
		name:      name,
		createdAt: time.Now(),
	}, nil
}

// геттеры
func (c *City) ID() int64            { return c.id }
func (c *City) Name() string         { return c.name }
func (c *City) CreatedAt() time.Time { return c.createdAt }

// SetID нужен для repository когда получаем из бд и нужно установить айди
func (c *City) SetID(id int64) { c.id = id }

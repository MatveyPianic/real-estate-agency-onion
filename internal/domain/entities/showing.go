package entities

import (
	"real-estate-agency-onion/internal/domain/enums"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"time"
)

type Showing struct {
	id         int64
	propertyID int64
	agentID    int64
	startsAt   time.Time
	endsAt     time.Time
	status     enums.ShowingStatus
	createdAt  time.Time
	deletedAt  *time.Time
}

func NewShowing(
	propertyID int64,
	agentID int64,
	startsAt time.Time,
	endsAt time.Time,
) (*Showing, error) {
	// валидация
	if propertyID <= 0 || agentID <= 0 {
		return nil, domainerrors.ErrInvalidInput
	}
	// ВАЖНО: проверка времени
	if endsAt.Before(startsAt) || endsAt.Equal(startsAt) {
		return nil, domainerrors.ErrInvalidInput
	}

	return &Showing{
		propertyID: propertyID,
		agentID:    agentID,
		startsAt:   startsAt,
		endsAt:     endsAt,
		status:     enums.ShowingStatusScheduled,
	}, nil
}

// геттеры
func (s *Showing) ID() int64                   { return s.id }
func (s *Showing) PropertyID() int64           { return s.propertyID }
func (s *Showing) AgentID() int64              { return s.agentID }
func (s *Showing) StartsAt() time.Time         { return s.startsAt }
func (s *Showing) EndsAt() time.Time           { return s.endsAt }
func (s *Showing) Status() enums.ShowingStatus { return s.status }
func (s *Showing) CreatedAt() time.Time        { return s.createdAt }
func (s *Showing) DeletedAt() *time.Time       { return s.deletedAt }

// вспомогательные методы
func (s *Showing) IsDeleted() bool {
	return s.deletedAt != nil
}

func (s *Showing) Duration() time.Duration {
	return s.endsAt.Sub(s.startsAt)
}

// бизнес логика
func (s *Showing) Cancel() error {
	if s.status == enums.ShowingStatusDone {
		return domainerrors.ErrForbidden
	}
	if s.status == enums.ShowingStatusCanceled {
		return domainerrors.ErrForbidden
	}
	s.status = enums.ShowingStatusCanceled
	return nil
}

func (s *Showing) MarkIsDone() error {
	if s.status != enums.ShowingStatusScheduled {
		return domainerrors.ErrForbidden
	}
	s.status = enums.ShowingStatusDone
	return nil
}

func (s *Showing) Reschedule(newStartsAt, newEndsAt time.Time) error {
	if s.status != enums.ShowingStatusScheduled {
		return domainerrors.ErrForbidden
	}
	if newEndsAt.Before(newStartsAt) || newEndsAt.Equal(newStartsAt) {
		return domainerrors.ErrInvalidInput
	}
	s.startsAt = newStartsAt
	s.endsAt = newEndsAt
	return nil
}

func (s *Showing) SoftDelete() error {
	if s.IsDeleted() {
		return domainerrors.ErrForbidden
	}
	now := time.Now()
	s.deletedAt = &now
	return nil
}

func (s *Showing) SetID(id int64) {
	s.id = id
}

func (s *Showing) SetCreatedAt(t time.Time) {
	s.createdAt = t
}

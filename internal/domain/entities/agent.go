package entities

import (
	"strings"
	"time"

	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/valueobjects"
)

type Agent struct {
	id         int64
	userID     *int64
	firstName  string
	lastName   string
	middleName *string
	phone      valueobjects.PhoneNumber
	telegram   *string
	whatsapp   *string
	photoPath  *string
	isActive   bool
	createdAt  time.Time
	deletedAt  *time.Time
}

// NewAgent - конструктор для создания нового агента
func NewAgent(
	firstName string, // НЕ указатель - обязательное поле
	lastName string, // НЕ указатель - обязательное поле
	middleName *string, // указатель - nullable
	phone valueobjects.PhoneNumber,
	telegram *string,
	whatsapp *string,
) (*Agent, error) {
	// Валидация и trim обязательных полей
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	if firstName == "" {
		return nil, domainerrors.ErrInvalidInput
	}
	if lastName == "" {
		return nil, domainerrors.ErrInvalidInput
	}

	// middleName может быть nil - это нормально
	if middleName != nil {
		trimmed := strings.TrimSpace(*middleName)
		middleName = &trimmed
	}

	return &Agent{
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
		phone:      phone,
		telegram:   telegram,
		whatsapp:   whatsapp,
		isActive:   true,
	}, nil
}

func RestoreAgent(
	id int64,
	userID *int64,
	firstName string,
	lastName string,
	middleName *string,
	phone valueobjects.PhoneNumber,
	telegram *string,
	whatsapp *string,
	photoPath *string,
	isActive bool,
	createdAt time.Time,
	deletedAt *time.Time,
) (*Agent, error) {
	agent, err := NewAgent(
		firstName,
		lastName,
		middleName,
		phone,
		telegram,
		whatsapp,
	)
	if err != nil {
		return nil, err
	}

	agent.id = id
	agent.userID = userID
	agent.isActive = isActive
	agent.photoPath = photoPath
	agent.createdAt = createdAt
	agent.deletedAt = deletedAt

	return agent, nil
}

// Геттеры
func (a *Agent) ID() int64                       { return a.id }
func (a *Agent) UserID() *int64                  { return a.userID }
func (a *Agent) FirstName() string               { return a.firstName }
func (a *Agent) LastName() string                { return a.lastName }
func (a *Agent) MiddleName() *string             { return a.middleName }
func (a *Agent) Phone() valueobjects.PhoneNumber { return a.phone }
func (a *Agent) Telegram() *string               { return a.telegram }
func (a *Agent) Whatsapp() *string               { return a.whatsapp }
func (a *Agent) PhotoPath() *string              { return a.photoPath }
func (a *Agent) IsActive() bool                  { return a.isActive }
func (a *Agent) CreatedAt() time.Time            { return a.createdAt }
func (a *Agent) DeletedAt() *time.Time           { return a.deletedAt }

// FullName - полное имя агента
func (a *Agent) FullName() string {
	if a.middleName != nil && *a.middleName != "" {
		return a.lastName + " " + a.firstName + " " + *a.middleName
	}
	return a.lastName + " " + a.firstName
}

// IsDeleted - проверка удаления
func (a *Agent) IsDeleted() bool {
	return a.deletedAt != nil
}

// AssignUser - привязать пользователя к агенту
func (a *Agent) AssignUser(userID int64) error {
	if userID <= 0 {
		return domainerrors.ErrInvalidInput
	}
	if a.userID != nil {
		return domainerrors.ErrAlreadyExists
	}
	a.userID = &userID
	return nil
}

// UnassignUser - отвязать пользователя
func (a *Agent) UnassignUser() {
	a.userID = nil
}

// Activate - активировать агента
func (a *Agent) Activate() error {
	if a.IsDeleted() {
		return domainerrors.ErrForbidden
	}
	a.isActive = true
	return nil
}

// Deactivate - деактивировать агента
func (a *Agent) Deactivate() {
	a.isActive = false
}

// UpdateName - обновить имя агента
func (a *Agent) UpdateName(firstName, lastName string, middleName *string) error {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	if firstName == "" || lastName == "" {
		return domainerrors.ErrInvalidInput
	}

	a.firstName = firstName
	a.lastName = lastName

	if middleName != nil {
		trimmed := strings.TrimSpace(*middleName)
		a.middleName = &trimmed
	} else {
		a.middleName = nil
	}
	return nil
}

// SoftDelete - мягкое удаление агента
func (a *Agent) SoftDelete() error {
	if a.IsDeleted() {
		return domainerrors.ErrForbidden
	}
	now := time.Now()
	a.deletedAt = &now
	a.isActive = false
	return nil
}

// SetID - установить ID (для repository)
func (a *Agent) SetID(id int64) {
	a.id = id
}

// SetCreatedAt - установить время создания (для repository)
func (a *Agent) SetCreatedAt(t time.Time) {
	a.createdAt = t
}

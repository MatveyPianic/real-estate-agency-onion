package entities

import (
	"strings"
	"time"

	domainerrors "real-estate-agency-onion/internal/domain/errors"
)

type User struct {
	id            int64
	email         string
	passwordHash  string
	roles         []string // роли админ, агент и менеджер
	isActive      bool
	emailVerified bool
	lastLoginAt   *time.Time
	createdAt     time.Time
	updatedAt     time.Time
}

// конструктор
func NewUser(
	email string,
	passwordHash string, // пароль уже захеширован
	roles []string,
) (*User, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return nil, domainerrors.ErrInvalidInput
	}
	if passwordHash == "" {
		return nil, domainerrors.ErrInvalidInput
	}
	if len(roles) == 0 {
		return nil, domainerrors.ErrInvalidInput
	}

	return &User{
		email:         email,
		passwordHash:  passwordHash,
		roles:         roles,
		isActive:      true,
		emailVerified: false,
	}, nil
}

// Геттеры
func (u *User) ID() int64               { return u.id }
func (u *User) Email() string           { return u.email }
func (u *User) PasswordHash() string    { return u.passwordHash }
func (u *User) Roles() []string         { return u.roles }
func (u *User) IsActive() bool          { return u.isActive }
func (u *User) EmailVerified() bool     { return u.emailVerified }
func (u *User) LastLoginAt() *time.Time { return u.lastLoginAt }
func (u *User) CreatedAt() time.Time    { return u.createdAt }
func (u *User) UpdatedAt() time.Time    { return u.updatedAt }

// HasRole проверяет наличие роли
func (u *User) HasRole(role string) bool {
	for _, r := range u.roles {
		if r == role {
			return true
		}
	}
	return false
}

// проверка на админа
func (u *User) IsAdmin() bool {
	return u.HasRole("admin")
}

// проверка на агента
func (u *User) IsAgent() bool {
	return u.HasRole("agent")
}

// проверка на менеджера
func (u *User) IsManager() bool {
	return u.HasRole("manager")
}

// Бизнес-методы

// AddRole добавляет роль пользователю
func (u *User) AddRole(role string) error {
	role = strings.TrimSpace(role)
	if role == "" {
		return domainerrors.ErrInvalidInput
	}
	if u.HasRole(role) {
		return domainerrors.ErrAlreadyExists
	}
	u.roles = append(u.roles, role)
	u.updatedAt = time.Now()
	return nil
}

// RemoveRole удаляет роль у пользователя
func (u *User) RemoveRole(role string) error {
	if !u.HasRole(role) {
		return domainerrors.ErrNotFound
	}

	newRoles := make([]string, 0, len(u.roles)-1)
	for _, r := range u.roles {
		if r != role {
			newRoles = append(newRoles, r)
		}
	}
	u.roles = newRoles
	u.updatedAt = time.Now()
	return nil
}

// Activate активирует пользователя
func (u *User) Activate() {
	u.isActive = true
	u.updatedAt = time.Now()
}

// Deactivate деактивирует пользователя
func (u *User) Deactivate() {
	u.isActive = false
	u.updatedAt = time.Now()
}

// VerifyEmail подтверждает email
func (u *User) VerifyEmail() {
	u.emailVerified = true
	u.updatedAt = time.Now()
}

// UpdateLastLogin обновляет время последнего входа
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.lastLoginAt = &now
	u.updatedAt = now
}

// ChangePassword меняет пароль (принимает уже хешированный!)
func (u *User) ChangePassword(newPasswordHash string) error {
	if newPasswordHash == "" {
		return domainerrors.ErrInvalidInput
	}
	u.passwordHash = newPasswordHash
	u.updatedAt = time.Now()
	return nil
}

// Сеттеры для repository
func (u *User) SetID(id int64) {
	u.id = id
}

func (u *User) SetCreatedAt(t time.Time) {
	u.createdAt = t
}

func (u *User) SetUpdatedAt(t time.Time) {
	u.updatedAt = t
}

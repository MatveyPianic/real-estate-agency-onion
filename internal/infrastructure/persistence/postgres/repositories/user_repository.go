package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"real-estate-agency-onion/internal/domain/entities"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/infrastructure/persistence/postgres/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	// TODO: позже
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	// TODO: позже
	return nil, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, email, password_hash, is_active, email_verified, last_login_at, created_at, updated_at 
	          FROM users WHERE email = $1`
	var model models.UserModel
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&model.ID, &model.Email, &model.PasswordHash, &model.IsActive,
		&model.EmailVerified, &model.LastLoginAt, &model.CreatedAt, &model.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domainerrors.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("user repository: get by email: %w", err)
	}
	return modelToUser(&model), nil
}

func (r *UserRepository) List(ctx context.Context, filters interface{}, pagination interface{}) ([]*entities.User, int64, error) {
	return nil, 0, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func modelToUser(m *models.UserModel) *entities.User {
	// упрощённая версия — роли пока не загружаем
	u := &entities.User{}
	u.SetID(m.ID)
	// остальные поля через сеттеры (если есть) или напрямую через Restore
	// для простоты пока возвращаем заглушку
	return u
}
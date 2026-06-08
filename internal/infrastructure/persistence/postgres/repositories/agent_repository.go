package repositories

import (
    "context"
    "database/sql"
    "fmt"
    "strings"

    "real-estate-agency-onion/internal/domain/entities"
    domainerrors "real-estate-agency-onion/internal/domain/errors"
    domainrepos "real-estate-agency-onion/internal/domain/repositories"
    "real-estate-agency-onion/internal/domain/valueobjects"
    "real-estate-agency-onion/internal/infrastructure/persistence/postgres/mappers"
    "real-estate-agency-onion/internal/infrastructure/persistence/postgres/models"
)

type AgentRepository struct {
    db *sql.DB
}

func NewAgentRepository(db *sql.DB) *AgentRepository {
    return &AgentRepository{db: db}
}

func (r *AgentRepository) Create(ctx context.Context, agent *entities.Agent) error {
    model := mappers.AgentDomainToModel(agent)
    query := `INSERT INTO agents (user_id, last_name, first_name, middle_name, phone, telegram, whatsapp, is_active, photo_path)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
              RETURNING id, created_at`
    err := r.db.QueryRowContext(ctx, query,
        model.UserID, model.LastName, model.FirstName, model.MiddleName,
        model.Phone, model.Telegram, model.Whatsapp, model.IsActive, model.PhotoPath,
    ).Scan(&model.ID, &model.CreatedAt)
    if err != nil {
        return fmt.Errorf("agent repository: create: %w", err)
    }
    agent.SetID(model.ID)
    agent.SetCreatedAt(model.CreatedAt)
    return nil
}

func (r *AgentRepository) GetByID(ctx context.Context, id int64) (*entities.Agent, error) {
    query := `SELECT id, user_id, first_name, last_name, middle_name, phone, telegram, whatsapp, is_active, photo_path, created_at, deleted_at
              FROM agents WHERE id = $1 AND deleted_at IS NULL`
    var model models.AgentModel
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &model.ID, &model.UserID, &model.FirstName, &model.LastName, &model.MiddleName,
        &model.Phone, &model.Telegram, &model.Whatsapp, &model.IsActive, &model.PhotoPath,
        &model.CreatedAt, &model.DeletedAt,
    )
    if err == sql.ErrNoRows {
        return nil, domainerrors.ErrNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("agent repository: get by id: %w", err)
    }
    return mappers.AgentModelToDomain(model)
}

func (r *AgentRepository) List(ctx context.Context, filters domainrepos.AgentFilters, pagination domainrepos.Pagination) ([]*entities.Agent, int64, error) {
    var conditions []string
    var args []interface{}
    argIdx := 1

    conditions = append(conditions, "deleted_at IS NULL")

    if filters.IsActive != nil {
        conditions = append(conditions, fmt.Sprintf("is_active = $%d", argIdx))
        args = append(args, *filters.IsActive)
        argIdx++
    }

    if filters.HasUser != nil {
        if *filters.HasUser {
            conditions = append(conditions, "user_id IS NOT NULL")
        } else {
            conditions = append(conditions, "user_id IS NULL")
        }
    }

    whereClause := strings.Join(conditions, " AND ")

    // count total
    var total int64
    countQuery := fmt.Sprintf("SELECT COUNT(*) FROM agents WHERE %s", whereClause)
    err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("agent repository: list count: %w", err)
    }

    // select page
    selectQuery := fmt.Sprintf(
        `SELECT id, user_id, first_name, last_name, middle_name, phone, telegram, whatsapp, is_active, photo_path, created_at, deleted_at
         FROM agents WHERE %s ORDER BY id LIMIT $%d OFFSET $%d`,
        whereClause, argIdx, argIdx+1,
    )
    args = append(args, pagination.Limit, pagination.Offset)

    rows, err := r.db.QueryContext(ctx, selectQuery, args...)
    if err != nil {
        return nil, 0, fmt.Errorf("agent repository: list query: %w", err)
    }
    defer rows.Close()

    agents := make([]*entities.Agent, 0)
    for rows.Next() {
        var model models.AgentModel
        err := rows.Scan(
            &model.ID, &model.UserID, &model.FirstName, &model.LastName, &model.MiddleName,
            &model.Phone, &model.Telegram, &model.Whatsapp, &model.IsActive, &model.PhotoPath,
            &model.CreatedAt, &model.DeletedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("agent repository: list scan: %w", err)
        }
        agent, err := mappers.AgentModelToDomain(model)
        if err != nil {
            return nil, 0, fmt.Errorf("agent repository: list map: %w", err)
        }
        agents = append(agents, agent)
    }
    if err := rows.Err(); err != nil {
        return nil, 0, fmt.Errorf("agent repository: list rows: %w", err)
    }

    return agents, total, nil
}

func (r *AgentRepository) Update(ctx context.Context, agent *entities.Agent) error {
    model := mappers.AgentDomainToModel(agent)
    query := `UPDATE agents
              SET user_id = $1, last_name = $2, first_name = $3, middle_name = $4,
                  phone = $5, telegram = $6, whatsapp = $7, is_active = $8, photo_path = $9
              WHERE id = $10 AND deleted_at IS NULL`
    result, err := r.db.ExecContext(ctx, query,
        model.UserID, model.LastName, model.FirstName, model.MiddleName,
        model.Phone, model.Telegram, model.Whatsapp, model.IsActive, model.PhotoPath,
        model.ID,
    )
    if err != nil {
        return fmt.Errorf("agent repository: update: %w", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("agent repository: update rows affected: %w", err)
    }
    if rowsAffected == 0 {
        return domainerrors.ErrNotFound
    }
    return nil
}

func (r *AgentRepository) SoftDelete(ctx context.Context, id int64) error {
    query := `UPDATE agents SET deleted_at = NOW(), is_active = false WHERE id = $1 AND deleted_at IS NULL`
    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("agent repository: soft delete: %w", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("agent repository: soft delete rows affected: %w", err)
    }
    if rowsAffected == 0 {
        return domainerrors.ErrNotFound
    }
    return nil
}

func (r *AgentRepository) GetByUserID(ctx context.Context, userID int64) (*entities.Agent, error) {
    query := `SELECT id, user_id, first_name, last_name, middle_name, phone, telegram, whatsapp, is_active, photo_path, created_at, deleted_at
              FROM agents WHERE user_id = $1 AND deleted_at IS NULL`
    var model models.AgentModel
    err := r.db.QueryRowContext(ctx, query, userID).Scan(
        &model.ID, &model.UserID, &model.FirstName, &model.LastName, &model.MiddleName,
        &model.Phone, &model.Telegram, &model.Whatsapp, &model.IsActive, &model.PhotoPath,
        &model.CreatedAt, &model.DeletedAt,
    )
    if err == sql.ErrNoRows {
        return nil, domainerrors.ErrNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("agent repository: get by user id: %w", err)
    }
    return mappers.AgentModelToDomain(model)
}

func (r *AgentRepository) GetByPhone(ctx context.Context, phone valueobjects.PhoneNumber) (*entities.Agent, error) {
    query := `SELECT id, user_id, first_name, last_name, middle_name, phone, telegram, whatsapp, is_active, photo_path, created_at, deleted_at
              FROM agents WHERE phone = $1 AND deleted_at IS NULL`
    var model models.AgentModel
    err := r.db.QueryRowContext(ctx, query, phone.Value()).Scan(
        &model.ID, &model.UserID, &model.FirstName, &model.LastName, &model.MiddleName,
        &model.Phone, &model.Telegram, &model.Whatsapp, &model.IsActive, &model.PhotoPath,
        &model.CreatedAt, &model.DeletedAt,
    )
    if err == sql.ErrNoRows {
        return nil, domainerrors.ErrNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("agent repository: get by phone: %w", err)
    }
    return mappers.AgentModelToDomain(model)
}
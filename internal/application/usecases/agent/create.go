package agent

import (
	"context"
	"errors"
	"strings"

	"real-estate-agency-onion/internal/application/dto"
	"real-estate-agency-onion/internal/domain/entities"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/repositories"
	"real-estate-agency-onion/internal/domain/valueobjects"
)

// type CreateInput struct {
//     FirstName  string
//     LastName   string
//     MiddleName *string
//     Phone      string
//     Telegram   *string
//     Whatsapp   *string
// }

// type CreateOutput struct {
//     ID        int64
//     FullName  string
//     Phone     string
//     IsActive  bool
// }

type CreateUseCase struct {
	agentRepo repositories.AgentRepository
}

func NewCreateUseCase(agentRepo repositories.AgentRepository) *CreateUseCase {
	return &CreateUseCase{
		agentRepo: agentRepo,
	}
}

func (uc *CreateUseCase) Execute(ctx context.Context, in dto.CreateAgentInput) (dto.AgentOutput, error) {
	// базовая нормализация входа
	in.FirstName = strings.TrimSpace(in.FirstName)
	in.LastName = strings.TrimSpace(in.LastName)

	// VO для телефона
	phoneVO, err := valueobjects.NewPhone(in.Phone)
	if err != nil {
		return dto.AgentOutput{}, domainerrors.ErrInvalidInput
	}

	// проверка уникальности телефона
	existing, err := uc.agentRepo.GetByPhone(ctx, phoneVO)
	if err != nil && !errors.Is(err, domainerrors.ErrNotFound) {
		return dto.AgentOutput{}, err
	}
	if existing != nil {
		return dto.AgentOutput{}, domainerrors.ErrAlreadyExists
	}

	// создание доменной сущности
	agentEntity, err := entities.NewAgent(
		in.FirstName,
		in.LastName,
		normalizeOptional(in.MiddleName),
		phoneVO,
		normalizeOptional(in.Telegram),
		normalizeOptional(in.Whatsapp),
	)
	if err != nil {
		return dto.AgentOutput{}, err
	}

	// сохранение
	if err := uc.agentRepo.Create(ctx, agentEntity); err != nil {
		return dto.AgentOutput{}, err
	}

	// возврат результата
	return toAgentOutput(agentEntity), nil
}

func normalizeOptional(v *string) *string {
	if v == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*v)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

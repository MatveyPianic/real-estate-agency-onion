package agent

import (
	"context"
	"real-estate-agency-onion/internal/application/dto"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/repositories"
)

type UpdateUseCase struct {
	agentRepo repositories.AgentRepository
}

func NewUpdateUseCase(agentRepo repositories.AgentRepository) *UpdateUseCase {
	return &UpdateUseCase{
		agentRepo: agentRepo,
	}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, input dto.UpdateAgentInput) error {
	if input.ID <= 0 {
		return domainerrors.ErrInvalidInput
	}

	agent, err := uc.agentRepo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	firstName := agent.FirstName()
	if input.FirstName != nil {
		firstName = *input.FirstName
	}

	lastName := agent.LastName()
	if input.LastName != nil {
		lastName = *input.LastName
	}

	middleName := agent.MiddleName()
	if input.MiddleName != nil {
		middleName = normalizeOptional(input.MiddleName)
	}

	err = agent.UpdateName(firstName, lastName, middleName)
	if err != nil {
		return err
	}

	return uc.agentRepo.Update(ctx, agent)
}

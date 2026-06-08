package agent

import (
	"context"
	"real-estate-agency-onion/internal/application/dto"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/repositories"
)

type GetAgentByIDUseCase struct {
	agentRepo repositories.AgentRepository
}

func NewGetAgentByIDUseCase(agentRepo repositories.AgentRepository) *GetAgentByIDUseCase {
	return &GetAgentByIDUseCase{
		agentRepo: agentRepo,
	}
}

func (uc *GetAgentByIDUseCase) Execute(ctx context.Context, input dto.GetAgentByIDInput) (*dto.AgentOutput, error) {
	if input.ID <= 0 {
		return nil, domainerrors.ErrInvalidInput
	}

	agent, err := uc.agentRepo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	out := toAgentOutput(agent)
	return &out, nil
}

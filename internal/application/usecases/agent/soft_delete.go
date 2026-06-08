package agent

import (
	"context"
	"real-estate-agency-onion/internal/application/dto"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/repositories"
)

type SoftDeleteUseCase struct {
	agentRepo repositories.AgentRepository
}

func NewSoftDeleteUseCase(agentRepo repositories.AgentRepository) *SoftDeleteUseCase {
	return &SoftDeleteUseCase{
		agentRepo: agentRepo,
	}
}

func (uc *SoftDeleteUseCase) Execute(ctx context.Context, in dto.SoftDeleteAgentInput) error {
	if in.ID <= 0 {
		return domainerrors.ErrInvalidInput
	}

	agent, err := uc.agentRepo.GetByID(ctx, in.ID)
	if err != nil {
		return err
	}

	if err := agent.SoftDelete(); err != nil {
		return err
	}

	return uc.agentRepo.Update(ctx, agent)
}

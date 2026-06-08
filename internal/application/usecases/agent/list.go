package agent

import (
	"context"
	"real-estate-agency-onion/internal/application/dto"
	"real-estate-agency-onion/internal/domain/repositories"
)

type ListUseCase struct {
	agentRepo repositories.AgentRepository
}

func NewListUseCase(agentRepo repositories.AgentRepository) *ListUseCase {
	return &ListUseCase{
		agentRepo: agentRepo,
	}
}

func (uc *ListUseCase) Execute(ctx context.Context, in dto.ListAgentsInput) (dto.ListAgentsOutput, error) {
	limit := in.Limit
	offset := in.Offset

	if limit <= 0 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	filters := repositories.AgentFilters{
		IsActive: in.IsActive,
		HasUser:  in.HasUser,
	}

	pagination := repositories.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	agents, total, err := uc.agentRepo.List(ctx, filters, pagination)
	if err != nil {
		return dto.ListAgentsOutput{}, err
	}

	items := make([]dto.AgentOutput, 0, len(agents))
	for _, agent := range agents {
		items = append(items, toAgentOutput(agent))
	}
	return dto.ListAgentsOutput{
		Items: items,
		Total: total,
	}, nil
}

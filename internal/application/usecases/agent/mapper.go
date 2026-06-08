package agent

import (
	"real-estate-agency-onion/internal/application/dto"
	"real-estate-agency-onion/internal/domain/entities"
)

func toAgentOutput(agent *entities.Agent) dto.AgentOutput {
	return dto.AgentOutput{
		ID:        agent.ID(),
		UserID:    agent.UserID(),
		FullName:  agent.FullName(),
		Phone:     agent.Phone().Value(),
		IsActive:  agent.IsActive(),
		CreatedAt: agent.CreatedAt(),
	}
}

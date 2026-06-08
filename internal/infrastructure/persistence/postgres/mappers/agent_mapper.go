package mappers

import (
	"real-estate-agency-onion/internal/domain/entities"
	"real-estate-agency-onion/internal/domain/valueobjects"
	"real-estate-agency-onion/internal/infrastructure/persistence/postgres/models"
)

func AgentModelToDomain(model models.AgentModel) (*entities.Agent, error) {
	phone, err := valueobjects.NewPhone(model.Phone)
	if err != nil {
		return nil, err
	}

	return entities.RestoreAgent(
		model.ID,
		model.UserID,
		model.FirstName,
		model.LastName,
		model.MiddleName,
		phone,
		model.Telegram,
		model.Whatsapp,
		model.PhotoPath,
		model.IsActive,
		model.CreatedAt,
		model.DeletedAt,
	)
}

func AgentDomainToModel(agent *entities.Agent) models.AgentModel {
	return models.AgentModel{
		ID:         agent.ID(),
		UserID:     agent.UserID(),
		FirstName:  agent.FirstName(),
		LastName:   agent.LastName(),
		MiddleName: agent.MiddleName(),
		Phone:      agent.Phone().Value(),
		Telegram:   agent.Telegram(),
		Whatsapp:   agent.Whatsapp(),
		PhotoPath:  agent.PhotoPath(),
		IsActive:   agent.IsActive(),
		CreatedAt:  agent.CreatedAt(),
		DeletedAt:  agent.DeletedAt(),
	}
}

package dto

import "time"

type CreateAgentInput struct {
	FirstName  string
	LastName   string
	MiddleName *string
	Phone      string
	Telegram   *string
	Whatsapp   *string
}

type AgentOutput struct {
	ID        int64
	UserID    *int64
	FullName  string
	Phone     string
	IsActive  bool
	CreatedAt time.Time
}

type ListAgentsInput struct {
	IsActive *bool
	HasUser  *bool
	Limit    int
	Offset   int
}

type ListAgentsOutput struct {
	Items []AgentOutput
	Total int64
}

type DeactivateAgentInput struct {
	ID int64
}

// type DeactivateAgentOutput struct {
// 	ID       int64
// 	IsActive bool
// }

type GetAgentByIDInput struct {
	ID int64
}

type UpdateAgentInput struct {
	ID         int64
	FirstName  *string
	LastName   *string
	MiddleName *string
}

type SoftDeleteAgentInput struct {
	ID int64
}

package responses

import "time"

type AgentResponse struct {
	ID        int64     `json:"id"`
	UserID    *int64    `json:"user_id,omitempty"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type AgentListResponse struct {
	Data  []AgentResponse `json:"data"`
	Total int64           `json:"total"`
}
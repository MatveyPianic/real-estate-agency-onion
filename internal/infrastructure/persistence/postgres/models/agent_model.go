package models

import "time"

type AgentModel struct {
	ID         int64
	UserID     *int64
	FirstName  string
	LastName   string
	MiddleName *string
	Phone      string
	Telegram   *string
	Whatsapp   *string
	PhotoPath  *string
	IsActive   bool
	CreatedAt  time.Time
	DeletedAt  *time.Time
}

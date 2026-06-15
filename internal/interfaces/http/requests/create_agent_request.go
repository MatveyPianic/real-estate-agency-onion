package requests

type CreateAgentRequest struct {
	FirstName  string  `json:"first_name"  binding:"required"`
	LastName   string  `json:"last_name"   binding:"required"`
	MiddleName *string `json:"middle_name"`
	Phone      string  `json:"phone"       binding:"required"`
	Telegram   *string `json:"telegram"`
	Whatsapp   *string `json:"whatsapp"`
}
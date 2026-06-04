// internal/domain/valueobjects/email.go
package valueobjects

import (
    "regexp"
    "strings"
    
    domainerrors "real-estate-agency-onion/internal/domain/errors"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Email struct {
    value string
}

func NewEmail(email string) (Email, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    
    if email == "" {
        return Email{}, domainerrors.ErrInvalidInput
    }
    
    if !emailRegex.MatchString(email) {
        return Email{}, domainerrors.ErrInvalidInput
    }
    
    return Email{value: email}, nil
}

func (e Email) Value() string {
    return e.value
}

func (e Email) Domain() string {
    parts := strings.Split(e.value, "@")
    if len(parts) == 2 {
        return parts[1]
    }
    return ""
}
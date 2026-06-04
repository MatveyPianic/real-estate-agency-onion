package valueobjects

import (
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"regexp"
)

type PhoneNumber struct {
	value string // приватное
}

func NewPhone(number string) (PhoneNumber, error) {
	// валидация
	matched, err := regexp.MatchString(`^\+?[1-9]\d{1,14}$`, number)
	if err != nil {
		return PhoneNumber{}, err
	}
	if !matched {
		return PhoneNumber{}, domainerrors.ErrInvalidInput
	}
	return PhoneNumber{value: number}, nil
}

func (v PhoneNumber) Value() string { return v.value } // только геттеры

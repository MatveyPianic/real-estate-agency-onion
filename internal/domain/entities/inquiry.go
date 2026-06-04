package entities

import (
	"real-estate-agency-onion/internal/domain/enums"
	domainerrors "real-estate-agency-onion/internal/domain/errors"
	"real-estate-agency-onion/internal/domain/valueobjects"
	"strings"
	"time"
)

type Inquiry struct {
	id                int64
	propertyID        int64
	clientName        string
	clientPhoneNumber valueobjects.PhoneNumber
	clientEmail       valueobjects.Email
	comment           string
	status            enums.InquiryStatus
	createdAt         time.Time
	deletedAt         *time.Time
}

func NewInquiry(propertyID int64,
	clientname *string,
	clientphonenumber valueobjects.PhoneNumber,
	clientemail valueobjects.Email,
	comment string,
	status enums.InquiryStatus,
) (*Inquiry, error) {
	if clientname == nil {
		return nil, domainerrors.ErrInvalidInput
	}

	clientName := strings.TrimSpace(*clientname)
	if clientName == "" {
		return nil, domainerrors.ErrInvalidInput
	}

	return &Inquiry{
		propertyID:        propertyID,
		clientName:        clientName,
		clientPhoneNumber: clientphonenumber,
		clientEmail:       clientemail,
		comment:           comment,
		status:            status,
	}, nil
}

// геттеры
func (i *Inquiry) ID() int64                                   { return i.id }
func (i *Inquiry) PropertyID() int64                           { return i.propertyID }
func (i *Inquiry) ClientName() string                          { return i.clientName }
func (i *Inquiry) ClientPhoneNumber() valueobjects.PhoneNumber { return i.clientPhoneNumber }
func (i *Inquiry) ClientEmail() valueobjects.Email             { return i.clientEmail }
func (i *Inquiry) Comment() string                             { return i.comment }
func (i *Inquiry) Status() enums.InquiryStatus                 { return i.status }
func (i *Inquiry) CreatedAt() time.Time                        { return i.createdAt }

// вспомогательные
func (i *Inquiry) IsDeleted() bool {
	return i.deletedAt != nil
}

// бизнес методы

func (i *Inquiry) MarkAsScheduled() error {
	if i.status != enums.InquiryStatusNew {
		return domainerrors.ErrForbidden
	}

	i.status = enums.InquiryStatusScheduled
	return nil
}

func (i *Inquiry) Close() error {
	if i.status == enums.InquiryStatusScheduled {
		return domainerrors.ErrForbidden
	}
	i.status = enums.InquiryStatusClosed
	return nil
}

func (i *Inquiry) SoftDelete() error {
	if i.IsDeleted() {
		return domainerrors.ErrForbidden
	}
	now := time.Now()
	i.deletedAt = &now
	return nil
}

func (i *Inquiry) SetID(id int64) {
	i.id = id
}

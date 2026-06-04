package entities

import domainerrors "real-estate-agency-onion/internal/domain/errors"

type PropertyPhoto struct {
	id         int64
	propertyID int64
	filepath   string
	isCover    bool
}

func NewPropertyPhoto(propertyID int64, filepath string) (*PropertyPhoto, error) {
	if propertyID <= 0 {
		return nil, domainerrors.ErrInvalidInput
	}
	if filepath == "" {
		return nil, domainerrors.ErrInvalidInput
	}
	return &PropertyPhoto{
		propertyID: propertyID,
		filepath:   filepath,
	}, nil
}

// геттеры
func (p *PropertyPhoto) ID() int64         { return p.id }
func (p *PropertyPhoto) PropertyID() int64 { return p.propertyID }
func (p *PropertyPhoto) Filepath() string  { return p.filepath }
func (p *PropertyPhoto) IsCover() bool     { return p.isCover }

// SetID
func (p *PropertyPhoto) SetID(id int64) { p.id = id }

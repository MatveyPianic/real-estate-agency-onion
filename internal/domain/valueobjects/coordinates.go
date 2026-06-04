package valueobjects

import domainerrors "real-estate-agency-onion/internal/domain/errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinates(lat, lon float64) (Coordinates, error) {
	// валидация
	if lat < -90 || lat > 90 {
		return Coordinates{}, domainerrors.ErrInvalidInput
	}
	if lon < -180 || lon > 180 {
		return Coordinates{}, domainerrors.ErrInvalidInput
	}
	return Coordinates{
		latitude:  lat,
		longitude: lon,
	}, nil
}

// геттеры
func (v Coordinates) Latitude() float64 { return v.latitude }

func (v Coordinates) Longitude() float64 { return v.longitude }

package enums

type PropertyType string

const (
	PropertyTypeApartment  PropertyType = "apartment"
	PropertyTypeHouse      PropertyType = "house"
	PropertyTypeGarage     PropertyType = "garage"
	PropertyTypeLand       PropertyType = "land"
	PropertyTypeCommercial PropertyType = "commercial"
)

func (t PropertyType) IsValid() bool {
	switch t {
	case PropertyTypeApartment,
		PropertyTypeHouse,
		PropertyTypeGarage,
		PropertyTypeLand,
		PropertyTypeCommercial:
		return true
	}
	return false
}

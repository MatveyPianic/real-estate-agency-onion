package enums

type DealType string

const (
	DealTypeSale DealType = "sale"
	DealTypeRent DealType = "rent"
)

func (t DealType) IsValid() bool {
	switch t {
	case DealTypeSale, DealTypeRent:
		return true
	}
	return false
}

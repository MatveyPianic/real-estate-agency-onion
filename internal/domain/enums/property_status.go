package enums

type PropertyStatus string

const (
	PropertyStatusDraft     PropertyStatus = "draft"
	PropertyStatusPublished PropertyStatus = "published"
	PropertyStatusArchived  PropertyStatus = "archived"
)

func (s PropertyStatus) IsValid() bool {
	switch s {
	case PropertyStatusDraft, PropertyStatusPublished, PropertyStatusArchived:
		return true
	}
	return false
}

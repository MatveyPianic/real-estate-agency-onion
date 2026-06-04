package enums

type InquiryStatus string

const (
	InquiryStatusNew       InquiryStatus = "new"
	InquiryStatusScheduled InquiryStatus = "scheduled"
	InquiryStatusClosed    InquiryStatus = "closed"
)

func (s InquiryStatus) IsValid() bool {
	switch s {
	case InquiryStatusNew, InquiryStatusScheduled, InquiryStatusClosed:
		return true
	}
	return false
}

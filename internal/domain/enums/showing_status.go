package enums

type ShowingStatus string

const (
	ShowingStatusScheduled ShowingStatus = "scheduled"
	ShowingStatusDone      ShowingStatus = "done"
	ShowingStatusCanceled  ShowingStatus = "canceled"
)

func (s ShowingStatus) IsValid() bool {
	switch s {
	case ShowingStatusScheduled, ShowingStatusDone, ShowingStatusCanceled:
		return true
	}

	return false
}

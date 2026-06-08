package interfaces

import "time"

type Clock interface {
    Now() time.Time
}
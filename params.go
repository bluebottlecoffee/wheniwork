package wheniwork

import (
	"time"
)

type ListShiftParams struct {
	Start      time.Time
	End        time.Time
	LocationId []string
}

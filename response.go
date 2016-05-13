package wheniwork

import (
	"time"
)

type LoginResponse struct {
	Login *LoginResponseLogin `json:"login"`
}

type LoginResponseLogin struct {
	Token string `json:"token"`
}

type ListShiftsResponse struct {
	Shifts []Shift `json:"shifts"`
}

type Shift struct {
	Id             int64         `json:"id"`
	AccountId      int64         `json:"id"`
	UserId         int64         `json:"user_id"`
	LocationId     int64         `json:"location_id"`
	PositionId     int64         `json:"position_id"`
	SiteId         int64         `json:"site_id"`
	StartTime      *rFC1123ZTime `json:"start_time"`
	EndTime        *rFC1123ZTime `json:"end_time"`
	BreakTime      float64       `json:"break_time"`
	Color          string        `json:"color"`
	Notes          string        `json:"notes"`
	Instances      int64         `json:"instances"`
	Published      bool          `json:"published"`
	PublishedDate  *rFC1123ZTime `json:"published_date"`
	NotifiedAt     *rFC1123ZTime `json:"notified_at"`
	CreatedAt      *rFC1123ZTime `json:"created_at"`
	UpdatedAt      *rFC1123ZTime `json:"updated_at"`
	Acknowledged   int64         `json:"acknowledged"`
	AcknowledgedAt *rFC1123ZTime `json:"acknowledged_at"`
	CreatorId      int64         `json:"creator_id"`
	IsOpen         bool          `json:"is_open"`
}

type rFC1123ZTime struct {
	Time time.Time
}

func (ct *rFC1123ZTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	if len(b) > 0 {
		ct.Time, err = time.Parse(time.RFC1123Z, string(b))

		if err != nil {
			panic(err)
		}
	}
	return
}

func (ct *rFC1123ZTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(time.RFC1123Z)), nil
}

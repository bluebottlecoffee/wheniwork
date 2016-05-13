package wheniwork

import (
	"time"
)

type loginResponse struct {
	Login *loginResponseLogin `json:"login"`
}

type loginResponseLogin struct {
	Token string `json:"token"`
}

type ListShiftsResponse struct {
	Start  *RFC1123ZTime `json:"start"`
	End    *RFC1123ZTime `json:"end"`
	Shifts []Shift       `json:"shifts"`
}

type Shift struct {
	Id             int64         `json:"id"`
	AccountId      int64         `json:"account_id"`
	UserId         int64         `json:"user_id"`
	LocationId     int64         `json:"location_id"`
	PositionId     int64         `json:"position_id"`
	SiteId         int64         `json:"site_id"`
	StartTime      *RFC1123ZTime `json:"start_time"`
	EndTime        *RFC1123ZTime `json:"end_time"`
	BreakTime      float64       `json:"break_time"`
	Color          string        `json:"color"`
	Notes          string        `json:"notes"`
	Instances      int64         `json:"instances"`
	Published      bool          `json:"published"`
	PublishedDate  *RFC1123ZTime `json:"published_date"`
	NotifiedAt     *RFC1123ZTime `json:"notified_at"`
	CreatedAt      *RFC1123ZTime `json:"created_at"`
	UpdatedAt      *RFC1123ZTime `json:"updated_at"`
	Acknowledged   int64         `json:"acknowledged"`
	AcknowledgedAt *RFC1123ZTime `json:"acknowledged_at"`
	CreatorId      int64         `json:"creator_id"`
	IsOpen         bool          `json:"is_open"`
}

type GetShiftResponse struct {
	Shift *Shift `json:"shift"`
}

type RFC1123ZTime struct {
	Time time.Time
}

func (ct *RFC1123ZTime) UnmarshalJSON(b []byte) (err error) {
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

func (ct *RFC1123ZTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(time.RFC1123Z)), nil
}

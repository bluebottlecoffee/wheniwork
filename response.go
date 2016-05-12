package wheniwork

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
	Id int64 `json:"id"`
}

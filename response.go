package wheniwork

type LoginResponse struct {
	Login *LoginResponseLogin `json:"login"`
}

type LoginResponseLogin struct {
	Token string `json:"token"`
}

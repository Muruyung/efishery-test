package entity

type LoginBody struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

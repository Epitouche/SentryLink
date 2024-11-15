package schemas

type LoginCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterCredentials struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

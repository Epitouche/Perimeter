package schemas

type LoginCredentials struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type RegisterCredentials struct {
	Email    string `form:"email"    json:"email"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

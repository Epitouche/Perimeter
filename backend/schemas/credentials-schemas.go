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

type AuthenticationURL struct {
	URL string `json:"authentication_url"`
}

type JWT struct {
	Token string `json:"token"`
}

type CodeCredentials struct {
	Code string `form:"code"          json:"code"          binding:"required"`
	// State string `form:"state" json:"state"`
}

type TokenCredentials struct {
	Token    string `form:"token" json:"token" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	// State string `form:"state" json:"state"`
}

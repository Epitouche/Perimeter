package schemas

// SpotifyTokenResponse represents the response from Github when a token is requested.
type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type SpotifyUserInfo struct {
	Login      string `json:"display_name"`
	Email     string `json:"email"`
}

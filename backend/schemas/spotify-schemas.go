package schemas

type SpotifyAction string

type SpotifyReaction string

const (
	PlayMusic SpotifyReaction = "PlayMusic"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type SpotifyUserInfo struct {
	Login string `json:"display_name"`
	Email string `json:"email"`
}

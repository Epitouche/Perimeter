package schemas

import (
	"errors"
)

type SpotifyAction string // SpotifyAction is a string type to represent the action to be performed.

const (
	MusicPlayed SpotifyAction = "MusicPlayed" // MusicPlayed is the action for playing music.
)

type SpotifyReaction string // SpotifyReaction is a string type to represent the reaction to be performed.

const (
	SkipNextMusic     SpotifyReaction = "SkipNextMusic"     // SkipNextMusic is the reaction to skip to the next music.
	SkipPreviousMusic SpotifyReaction = "SkipPreviousMusic" // SkipPreviousMusic is the reaction to skip to the previous music.
)

// SpotifyTokenResponse represents the response from Spotify's token endpoint.
// It contains the access token, scope, token type, expiration time, and refresh token.
type SpotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`  // The access token
	Scope        string `json:"scope"`         // The scope of the token
	TokenType    string `json:"token_type"`    // The type of the token
	ExpiresIn    int    `json:"expires_in"`    // The expiration time of the token in seconds
	RefreshToken string `json:"refresh_token"` // The refresh token
}

// SpotifyErrorResponse represents the structure of an error response from the Spotify API.
// It contains an embedded Error struct which holds the status code and error message.
type SpotifyErrorResponse struct {
	Error struct {
		Status  int    `json:"status"`  // The status code of the error
		Message string `json:"message"` // The error message
	} `json:"error"` // The error object
}

type SpotifyUserInfo struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	Uri     string `json:"uri"`
}

type SpotifyPlaybackResponse struct {
	IsPlaying bool `json:"is_playing"`
	Item      struct {
		Name    string `json:"name"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
	} `json:"item"`
}

type SpotifyActionMusicPlayedOption struct {
	Name string `json:"name"`
}

type SpotifyStorageVariable int

const (
	SpotifyStorageVariableInit  SpotifyStorageVariable = 0
	SpotifyStorageVariableTrue  SpotifyStorageVariable = 1
	SpotifyStorageVariableFalse SpotifyStorageVariable = 2
)

// Errors Messages.
var (
	ErrSpotifySecretNotSet   = errors.New("SPOTIFY_SECRET is not set")
	ErrSpotifyClientIdNotSet = errors.New("SPOTIFY_CLIENT_ID is not set")
)

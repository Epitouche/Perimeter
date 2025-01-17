package schemas

import (
	"errors"
)

type SpotifyAction string

const (
	MusicPlayed SpotifyAction = "MusicPlayed"
)

type SpotifyReaction string

const (
	SkipNextMusic     SpotifyReaction = "SkipNextMusic"
	SkipPreviousMusic SpotifyReaction = "SkipPreviousMusic"
)

type SpotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type SpotifyErrorResponse struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
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

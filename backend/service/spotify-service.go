package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type SpotifyService interface {
	// Service interface functions
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	SpotifyActionMusicPlayed(c chan string, option json.RawMessage, idArea uint64)
	// Reactions functions
	SpotifyReactionSkipNextMusic(option json.RawMessage, idArea uint64) string
	SpotifyReactionSkipPreviousMusic(option json.RawMessage, idArea uint64) string
}

type spotifyService struct {
	repository        repository.SpotifyRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

func NewSpotifyService(
	githubTokenRepository repository.SpotifyRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) SpotifyService {
	return &spotifyService{
		repository:        githubTokenRepository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Spotify,
			Description: "This service is a music service",
			Oauth:       true,
			Color:       "#1DC000",
			Icon:        "https://api.iconify.design/mdi:spotify.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *spotifyService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *spotifyService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	case string(schemas.MusicPlayed):
		return service.SpotifyActionMusicPlayed
	default:
		return nil
	}
}

func (service *spotifyService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	case string(schemas.SkipNextMusic):
		return service.SpotifyReactionSkipNextMusic
	case string(schemas.SkipPreviousMusic):
		return service.SpotifyReactionSkipPreviousMusic
	default:
		return nil
	}
}

func (service *spotifyService) GetServiceActionInfo() []schemas.Action {
	defaultValue := schemas.SpotifyActionMusicPlayedOption{
		Name: "",
	}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Spotify,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.MusicPlayed),
			Description:        "This action check if a music is played",
			Service:            service.serviceInfo,
			Option:             option,
			MinimumRefreshRate: 10,
		},
	}
}

func (service *spotifyService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := struct{}{}
	option, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Spotify,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.SkipNextMusic),
			Description: "This reaction will skip to the next music",
			Service:     service.serviceInfo,
			Option:      option,
		},
		{
			Name:        string(schemas.SkipPreviousMusic),
			Description: "This reaction will skip to the previous music",
			Service:     service.serviceInfo,
			Option:      option,
		},
	}
}

// Service specific functions

func (service *spotifyService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrSpotifyClientIdNotSet
	}

	clientSecret := os.Getenv("SPOTIFY_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrSpotifySecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/spotify"

	apiURL := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to create request because %w",
			err,
		)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to make request because %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		println("Status code", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("body: %+v\n", body)
		return schemas.Token{}, fmt.Errorf(
			"unable to get token because %v",
			resp.Status,
		)
	}

	var result schemas.SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	if result.AccessToken == "" {
		fmt.Printf("Token exchange failed. Response body: %v\n", resp.Body)
		return schemas.Token{}, schemas.ErrAccessTokenNotFoundInResponse
	}

	resp.Body.Close()

	token = schemas.Token{
		Token:        result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpireAt:     time.Now().Add(time.Duration(result.ExpiresIn) * time.Second),
	}

	return token, nil
}

func (service *spotifyService) GetUserInfo(accessToken string) (user schemas.User, err error) {
	ctx := context.Background()
	// Create a new HTTP request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.spotify.com/v1/me",
		nil,
	)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	println("accessToken", accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to make request because %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errorResponse := schemas.SpotifyErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return schemas.User{}, fmt.Errorf(
				"unable to decode error response because %w",
				err,
			)
		}

		resp.Body.Close()
		return schemas.User{}, fmt.Errorf(
			"unable to get user info because %v %v",
			errorResponse.Error.Status,
			errorResponse.Error.Message,
		)
	}

	result := schemas.SpotifyUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Username: result.DisplayName,
		Email:    result.Email,
	}

	return user, nil
}

func getSpotifyPlaybackResponse(token schemas.Token) (schemas.SpotifyPlaybackResponse, error) {
	apiURL := "https://api.spotify.com/v1/me/player"

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return schemas.SpotifyPlaybackResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return schemas.SpotifyPlaybackResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status code %d\n", resp.StatusCode)
		return schemas.SpotifyPlaybackResponse{}, err
	}

	var playbackResponse schemas.SpotifyPlaybackResponse
	err = json.NewDecoder(resp.Body).Decode(&playbackResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return schemas.SpotifyPlaybackResponse{}, err
	}

	return playbackResponse, nil
}

// Actions functions
func (service *spotifyService) SpotifyActionMusicPlayed(
	c chan string,
	option json.RawMessage,
	idArea uint64,
) {
	optionJSON := schemas.SpotifyActionMusicPlayedOption{}
	err := json.Unmarshal(option, &optionJSON)
	if err != nil {
		fmt.Println("Error unmarshalling option:", err)
		return
	}

	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil || token.Token == "" {
		fmt.Println("Error finding token or token not found")
		return
	}

	playbackResponse, err := getSpotifyPlaybackResponse(token)
	if err != nil {
		fmt.Println("Error getting playback response:", err)
		return
	}

	if playbackResponse.IsPlaying {
		artistNames := []string{}
		for _, artist := range playbackResponse.Item.Artists {
			artistNames = append(artistNames, artist.Name)
		}
		if strings.EqualFold(playbackResponse.Item.Name, optionJSON.Name) {
			message := fmt.Sprintf("Currently playing: %s by %s",
				playbackResponse.Item.Name,
				strings.Join(artistNames, ", "),
			)
			fmt.Println(message)
			c <- message
		} else {
			message := fmt.Sprintf("Currently playing: %s by %s, but expected: %s",
				playbackResponse.Item.Name,
				strings.Join(artistNames, ", "),
				optionJSON.Name,
			)
			fmt.Println(message)
		}
	} else {
		fmt.Println("No music is currently playing.")
	}

	if (area.Action.MinimumRefreshRate) > area.ActionRefreshRate {
		time.Sleep(time.Second * time.Duration(area.Action.MinimumRefreshRate))
	} else {
		time.Sleep(time.Second * time.Duration(area.ActionRefreshRate))
	}
}

// Reactions functions
func (service *spotifyService) SpotifyReactionSkipNextMusic(
	option json.RawMessage,
	idArea uint64,
) string {
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area:" + err.Error()
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token:" + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}
	apiURL := "https://api.spotify.com/v1/me/player/next"

	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		apiURL,
		bytes.NewBuffer([]byte("{}")),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request:" + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request:" + err.Error()
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	return "Response Status:" + resp.Status
}

func (service *spotifyService) SpotifyReactionSkipPreviousMusic(
	option json.RawMessage,
	idArea uint64,
) string {
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return "Error finding area:" + err.Error()
	}

	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return "Error finding token:" + err.Error()
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return "Error: Token not found"
	}
	apiURL := "https://api.spotify.com/v1/me/player/previous"

	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		apiURL,
		bytes.NewBuffer([]byte("{}")),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request:" + err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request:" + err.Error()
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	return "Response Status:" + resp.Status
}

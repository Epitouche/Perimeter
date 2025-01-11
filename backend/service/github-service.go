package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"area/repository"
	"area/schemas"
)

// Constructor

type GithubService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionbyName(name string) func(c chan string, option json.RawMessage, idArea uint64)
	FindReactionbyName(name string) func(option json.RawMessage, idArea uint64) string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	// Reactions functions
}

type githubService struct {
	repository        repository.GithubRepository
	serviceRepository repository.ServiceRepository
	areaRepository    repository.AreaRepository
	tokenRepository   repository.TokenRepository
	serviceInfo       schemas.Service
}

func NewGithubService(
	repository repository.GithubRepository,
	serviceRepository repository.ServiceRepository,
	areaRepository repository.AreaRepository,
	tokenRepository repository.TokenRepository,
) GithubService {
	return &githubService{
		repository:        repository,
		serviceRepository: serviceRepository,
		areaRepository:    areaRepository,
		tokenRepository:   tokenRepository,
		serviceInfo: schemas.Service{
			Name:        schemas.Github,
			Description: "This service is a code repository service",
			Oauth:       true,
			Color:       "#000000",
			Icon:        "https://api.iconify.design/mdi:github.svg?color=%23FFFFFF",
		},
	}
}

// Service interface functions

func (service *githubService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

func (service *githubService) GetServiceActionInfo() []schemas.Action {
	defaultValue := schemas.GithubActionUpdateCommitInRepo{
		RepoName: "",
	}
	actionUpdateCommitInRepo, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal timer option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Github,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:        string(schemas.UpdateCommitInRepo),
			Description: "This action trigger when a new commit is pushed to a repository",
			Service:     service.serviceInfo,
			Option:      actionUpdateCommitInRepo,
		},
		{
			Name:        string(schemas.UpdatePullRequestInRepo),
			Description: "This action trigger when a new pullrequest is open to a repository",
			Service:     service.serviceInfo,
			Option:      actionUpdateCommitInRepo,
		},
	}
}

func (service *githubService) GetServiceReactionInfo() []schemas.Reaction {
	return []schemas.Reaction{}
}

func (service *githubService) FindActionbyName(
	name string,
) func(c chan string, option json.RawMessage, idArea uint64) {
	switch name {
	case string(schemas.UpdateCommitInRepo):
		return service.GithubActionUpdateCommitInRepo
	case string(schemas.UpdatePullRequestInRepo):
		return service.GithubActionUpdatePullRequestInRepo
	default:
		return nil
	}
}

func (service *githubService) FindReactionbyName(
	name string,
) func(option json.RawMessage, idArea uint64) string {
	switch name {
	default:
		return nil
	}
}

// Service specific functions

func (service *githubService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return schemas.Token{}, schemas.ErrGithubClientIdNotSet
	}

	clientSecret := os.Getenv("GITHUB_SECRET")
	if clientSecret == "" {
		return schemas.Token{}, schemas.ErrGithubSecretNotSet
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		return schemas.Token{}, schemas.ErrBackendPortNotSet
	}

	redirectURI := "http://localhost:8081/services/github"

	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.Token{}, fmt.Errorf("unable to make request because %w", err)
	}

	var result schemas.GitHubTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.Token{}, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	resp.Body.Close()

	token = schemas.Token{
		Token: result.AccessToken,
		// RefreshToken:  result.RefreshToken,
		// ExpireAt: result.ExpiresIn,
	}
	return token, nil
}

func (service *githubService) GetUserEmail(accessToken string) (email string, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.github.com/user/emails",
		nil,
	)
	if err != nil {
		return email, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return email, fmt.Errorf("unable to make request because %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return email, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	result := []schemas.GithubUserEmail{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return email, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	for _, email := range result {
		if email.Primary {
			return email.Email, nil
		}
	}

	return email, fmt.Errorf("unable to find primary email")
}

func (service *githubService) GetUserInfoAccount(
	accessToken string,
) (user schemas.User, err error) {
	ctx := context.Background()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return user, fmt.Errorf("unable to create request because %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to make request because %w", err)
	}

	result := schemas.GithubUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	resp.Body.Close()

	user = schemas.User{
		Username: result.Login,
		Email:    result.Email,
	}
	return user, nil
}

func (service *githubService) GetUserInfo(accessToken string) (user schemas.User, err error) {
	user, err = service.GetUserInfoAccount(accessToken)
	if err != nil {
		return user, err
	}

	email, err := service.GetUserEmail(accessToken)
	if err != nil {
		return user, err
	}

	user = schemas.User{
		Username: user.Username,
		Email:    email,
	}

	return user, nil
}

func (service *githubService) IsCommitUpdate(
	commitList []schemas.GithubCommit,
	date time.Time,
) bool {
	for _, commit := range commitList {
		if commit.Commit.Author.Date.After(date) {
			return true
		}
	}

	return false
}

func (service *githubService) CommitList(
	userGithubToken string, repo string,
) (commitList []schemas.GithubCommit, err error) {
	ctx := context.Background()

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.github.com/repos/"+repo+"/commits",
		nil,
	)
	if err != nil {
		return commitList, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userGithubToken)
	req.Header.Set("Content-Type", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return commitList, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return commitList, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	// Decode the JSON response into the result struct
	err = json.NewDecoder(resp.Body).Decode(&commitList)
	if err != nil {
		return commitList, fmt.Errorf("unable to decode response: %w", err)
	}

	return commitList, nil
}

func (service *githubService) IsPullRequestUpdate(
	pullRequestList []schemas.GithubPullRequest,
	date time.Time,
) bool {
	for _, pullRequest := range pullRequestList {
		if pullRequest.CreatedAt.After(date) {
			return true
		}
	}

	return false
}

func (service *githubService) PullRequestList(
	userGithubToken string, repo string,
) (pullRequestList []schemas.GithubPullRequest, err error) {
	ctx := context.Background()

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.github.com/repos/"+repo+"/pulls",
		nil,
	)
	if err != nil {
		return pullRequestList, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userGithubToken)
	req.Header.Set("Content-Type", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return pullRequestList, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, _ := io.ReadAll(resp.Body)
		return pullRequestList, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	// Decode the JSON response into the result struct
	err = json.NewDecoder(resp.Body).Decode(&pullRequestList)
	if err != nil {
		return pullRequestList, fmt.Errorf("unable to decode response: %w", err)
	}

	return pullRequestList, nil
}

// Actions functions

func (service *githubService) GithubActionUpdateCommitInRepo(
	channel chan string,
	option json.RawMessage,
	idArea uint64,
) {
	// Find the area
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return
	}

	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return
	}

	databaseStored := schemas.GithubActionUpdateCommitInRepoStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.GithubActionUpdateCommitInRepoStorage{
				Time: time.Now(),
			}
			area.StorageVariable, err = json.Marshal(databaseStored)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
		}
	}

	if databaseStored.Time.IsZero() {
		println("initializing storage variable")
		databaseStored = schemas.GithubActionUpdateCommitInRepoStorage{
			Time: time.Now(),
		}
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
	}

	// Unmarshal the option
	optionJSON := schemas.GithubActionUpdateCommitInRepo{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	commitList, err := service.CommitList(token.Token, optionJSON.RepoName)
	if err != nil {
		fmt.Println(err)
	}

	if service.IsCommitUpdate(commitList, databaseStored.Time) {
		response := "new commit update in " + optionJSON.RepoName + " repository"
		databaseStored.Time = time.Now()
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
		println(response)
		channel <- response
	}

	time.Sleep(time.Minute)
}

func (service *githubService) GithubActionUpdatePullRequestInRepo(
	channel chan string,
	option json.RawMessage,
	idArea uint64,
) {
	// Find the area
	area, err := service.areaRepository.FindById(idArea)
	if err != nil {
		fmt.Println("Error finding area:", err)
		return
	}

	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Action.ServiceId,
	)
	if err != nil {
		fmt.Println("Error finding token:", err)
		return
	}
	if token.Token == "" {
		fmt.Println("Error: Token not found")
		return
	}

	databaseStored := schemas.GithubActionUpdatePullRequestInRepoStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.GithubActionUpdatePullRequestInRepoStorage{
				Time: time.Now(),
			}
			area.StorageVariable, err = json.Marshal(databaseStored)
			if err != nil {
				println("error marshalling storage variable: " + err.Error())
				return
			}
			err = service.areaRepository.Update(area)
			if err != nil {
				println("error updating area: " + err.Error())
				return
			}
		}
	}

	if databaseStored.Time.IsZero() {
		println("initializing storage variable")
		databaseStored = schemas.GithubActionUpdatePullRequestInRepoStorage{
			Time: time.Now(),
		}
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
	}

	// Unmarshal the option
	optionJSON := schemas.GithubActionUpdatePullRequestInRepo{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	pullRequestList, err := service.PullRequestList(token.Token, optionJSON.RepoName)
	if err != nil {
		fmt.Println(err)
	}

	if service.IsPullRequestUpdate(pullRequestList, databaseStored.Time) {
		response := "new commit update in " + optionJSON.RepoName + " repository"
		databaseStored.Time = time.Now()
		area.StorageVariable, err = json.Marshal(databaseStored)
		if err != nil {
			println("error marshalling storage variable: " + err.Error())
			return
		}
		err = service.areaRepository.Update(area)
		if err != nil {
			println("error updating area: " + err.Error())
			return
		}
		println(response)
		channel <- response
	}

	time.Sleep(time.Minute)
}

// Reactions functions

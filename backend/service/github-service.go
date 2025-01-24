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

// GithubService defines the interface for interacting with GitHub services.
// It includes methods for authentication, retrieving user information, and
// performing specific actions and reactions related to GitHub.
//
// Methods:
// GetServiceActionInfo returns a list of available actions for the service.
//
// GetServiceReactionInfo returns a list of available reactions for the service.
//
// FindActionByName finds an action by its name and returns a function that
// executes the action with the given parameters.
//
// FindReactionByName finds a reaction by its name and returns a function that
// executes the reaction with the given parameters.
//
// AuthGetServiceAccessToken exchanges a code for an access token.
//
// GetUserInfo retrieves user information using the provided access token.
//
// GithubActionUpdateCommitInRepo performs an action to update a commit in a repository.
//
// GithubReactionGetLatestCommitInRepo performs a reaction to get the latest commit in a repository.
type GithubService interface {
	// Service interface functions
	GetServiceActionInfo() []schemas.Action
	GetServiceReactionInfo() []schemas.Reaction
	FindActionByName(name string) func(c chan string, option json.RawMessage, area schemas.Area)
	FindReactionByName(name string) func(option json.RawMessage, area schemas.Area) string
	// Service specific functions
	AuthGetServiceAccessToken(code string) (token schemas.Token, err error)
	GetUserInfo(accessToken string) (user schemas.User, err error)
	// Actions functions
	GithubActionUpdateCommitInRepo(channel chan string, option json.RawMessage, area schemas.Area)
	// Reactions functions
	GithubReactionGetLatestCommitInRepo(option json.RawMessage, area schemas.Area) string
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

// GetServiceInfo returns the service information.
// It retrieves the service information from the githubService instance
// and returns it as a schemas.Service object.
func (service *githubService) GetServiceInfo() schemas.Service {
	return service.serviceInfo
}

// GetServiceActionInfo retrieves information about available GitHub service actions.
// It returns a slice of schemas.Action, each representing a specific GitHub action.
//
// The function performs the following steps:
// 1. Initializes a default GitHub action option with a placeholder repository name.
// 2. Marshals the default action option to JSON format.
// 3. Retrieves the service information from the service repository by the GitHub service name.
// 4. Constructs and returns a slice of schemas.Action with predefined actions, descriptions, and options.
//
// Returns:
//
//	[]schemas.Action: A slice containing information about GitHub actions.
//
// Errors:
//
//	Logs errors if marshalling the default action option or finding the service by name fails.
func (service *githubService) GetServiceActionInfo() []schemas.Action {
	defaultValue := schemas.GithubActionOption{
		RepoName: "OWNER/REPO",
	}
	actionOption, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal github option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Github,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Action{
		{
			Name:               string(schemas.UpdateCommitInRepo),
			Description:        "This action trigger when a new commit is pushed to a repository",
			Service:            service.serviceInfo,
			Option:             actionOption,
			MinimumRefreshRate: 10,
		},
		{
			Name:               string(schemas.UpdatePullRequestInRepo),
			Description:        "This action trigger when a new pullrequest is open to a repository",
			Service:            service.serviceInfo,
			Option:             actionOption,
			MinimumRefreshRate: 10,
		},
		{
			Name:               string(schemas.UpdateWorkflowRunInRepo),
			Description:        "This action trigger when a new workflow is run in a repository",
			Service:            service.serviceInfo,
			Option:             actionOption,
			MinimumRefreshRate: 10,
		},
	}
}

// GetServiceReactionInfo retrieves information about available reactions for the GitHub service.
// It initializes default options for the GitHub actions and fetches the service information from the repository.
// The function returns a slice of Reaction objects, each containing the name, description, service information, and options for the reaction.
func (service *githubService) GetServiceReactionInfo() []schemas.Reaction {
	defaultValue := schemas.GithubActionOption{
		RepoName: "OWNER/REPO",
	}
	actionOption, err := json.Marshal(defaultValue)
	if err != nil {
		println("error marshal github option: " + err.Error())
	}
	service.serviceInfo, err = service.serviceRepository.FindByName(
		schemas.Github,
	) // must update the serviceInfo
	if err != nil {
		println("error find service by name: " + err.Error())
	}
	return []schemas.Reaction{
		{
			Name:        string(schemas.GetLatestCommitInRepo),
			Description: "This reaction get the latest commit in a repository",
			Service:     service.serviceInfo,
			Option:      actionOption,
		},
		{
			Name:        string(schemas.GetLatestWorkflowRunInRepo),
			Description: "This reaction get the latest workflow run in a repository",
			Service:     service.serviceInfo,
			Option:      actionOption,
		},
	}
}

// FindActionByName returns a function that matches the given action name.
// The returned function takes a channel, a JSON raw message, and an area schema as parameters.
// If the action name does not match any predefined actions, it returns nil.
//
// Parameters:
//   - name: The name of the action to find.
//
// Returns:
//   - A function that matches the given action name, or nil if no match is found.
func (service *githubService) FindActionByName(
	name string,
) func(c chan string, option json.RawMessage, area schemas.Area) {
	switch name {
	case string(schemas.UpdateCommitInRepo):
		return service.GithubActionUpdateCommitInRepo
	case string(schemas.UpdatePullRequestInRepo):
		return service.GithubActionUpdatePullRequestInRepo
	case string(schemas.UpdateWorkflowRunInRepo):
		return service.GithubActionUpdateWorkflowRunInRepo
	default:
		return nil
	}
}

// FindReactionByName returns a function that matches the given reaction name.
// The returned function takes a json.RawMessage option and a schemas.Area as parameters,
// and returns a string. If the reaction name does not match any known reactions, it returns nil.
//
// Parameters:
//   - name: The name of the reaction to find.
//
// Returns:
//   - A function that takes a json.RawMessage option and a schemas.Area as parameters,
//     and returns a string, or nil if the reaction name does not match any known reactions.
func (service *githubService) FindReactionByName(
	name string,
) func(option json.RawMessage, area schemas.Area) string {
	switch name {
	case string(schemas.GetLatestCommitInRepo):
		return service.GithubReactionGetLatestCommitInRepo
	case string(schemas.GetLatestWorkflowRunInRepo):
		return service.GithubReactionGetLatestWorkflowRunInRepo
	default:
		return nil
	}
}

// Service specific functions

// AuthGetServiceAccessToken exchanges a GitHub authorization code for an access token.
// It determines the environment (production or development) and retrieves the appropriate
// client ID and client secret from environment variables. It then constructs a request
// to GitHub's OAuth API to obtain the access token.
//
// Parameters:
//   - code: The authorization code received from GitHub after user authorization.
//
// Returns:
//   - token: The access token received from GitHub.
//   - err: An error if the process fails at any step, including missing environment variables,
//     request creation, or response decoding.
func (service *githubService) AuthGetServiceAccessToken(
	code string,
) (token schemas.Token, err error) {
	isProd := os.Getenv("IS_PRODUCTION")
	if isProd == "" {
		return token, schemas.ErrIsProductionNotSet
	}

	clientID := ""
	if isProd == "true" {
		clientID = os.Getenv("GITHUB_PRODUCTION_CLIENT_ID")
		if clientID == "" {
			return token, schemas.ErrGithubProductionClientIdNotSet
		}
	} else {
		clientID = os.Getenv("GITHUB_CLIENT_ID")
		if clientID == "" {
			return token, schemas.ErrGithubClientIdNotSet
		}
	}

	clientSecret := ""
	if isProd == "true" {
		clientSecret = os.Getenv("GITHUB_PRODUCTION_SECRET")
		if clientSecret == "" {
			return token, schemas.ErrGithubProductionSecretNotSet
		}
	} else {
		clientSecret = os.Getenv("GITHUB_SECRET")
		if clientSecret == "" {
			return token, schemas.ErrGithubSecretNotSet
		}
	}

	redirectURI, err := getRedirectURI(service.serviceInfo.Name)
	if err != nil {
		return token, fmt.Errorf("unable to get redirect URI because %w", err)
	}

	apiURL := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
	if err != nil {
		return token, fmt.Errorf("unable to create request because %w", err)
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return token, fmt.Errorf("unable to make request because %w", err)
	}

	defer resp.Body.Close()

	var result schemas.GitHubTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return token, fmt.Errorf(
			"unable to decode response because %w",
			err,
		)
	}

	token = schemas.Token{
		Token: result.AccessToken,
		// RefreshToken:  result.RefreshToken,
		// ExpireAt: result.ExpiresIn,
	}
	return token, nil
}

// GetUserEmail retrieves the primary email address of a GitHub user using the provided access token.
// It sends a GET request to the GitHub API endpoint for user emails and parses the response to find
// the primary email address.
//
// Parameters:
//   - accessToken: A string containing the GitHub access token.
//
// Returns:
//   - email: A string containing the primary email address of the GitHub user.
//   - err: An error if the request fails, the response cannot be decoded, or no primary email is found.
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

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return email, fmt.Errorf("unable to read error response because %w", err)
		}
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

	for _, email := range result {
		if email.Primary {
			return email.Email, nil
		}
	}

	return email, fmt.Errorf("unable to find primary email")
}

// GetUserInfoAccount retrieves the GitHub user information using the provided access token.
// It sends a GET request to the GitHub API to fetch the user details.
//
// Parameters:
//   - accessToken: A string containing the GitHub access token.
//
// Returns:
//   - user: A schemas.User struct containing the username and email of the GitHub user.
//   - err: An error if the request fails or the response cannot be decoded.
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

	defer resp.Body.Close()

	result := schemas.GithubUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return schemas.User{}, fmt.Errorf("unable to decode response because %w", err)
	}

	user = schemas.User{
		Username: result.Login,
		Email:    result.Email,
	}
	return user, nil
}

// GetUserInfo retrieves the user information from GitHub using the provided access token.
// It first fetches the basic user account information and then retrieves the user's email.
// If any error occurs during these operations, it returns the error along with the user data retrieved so far.
//
// Parameters:
//   - accessToken: A string representing the GitHub access token.
//
// Returns:
//   - user: A schemas.User struct containing the user's username and email.
//   - err: An error if any occurs during the retrieval of user information.
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

// IsCommitUpdate checks if there is any commit in the provided commit list
// that has an author date after the specified date.
//
// Parameters:
//
//	commitList []schemas.GithubCommit - A list of GitHub commits to check.
//	date time.Time - The date to compare the commit author dates against.
//
// Returns:
//
//	bool - Returns true if there is at least one commit with an author date
//	       after the specified date, otherwise returns false.
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

// CommitList retrieves the list of commits from a specified GitHub repository.
//
// Parameters:
//   - userGithubToken: A string containing the GitHub token for authentication.
//   - repo: A string specifying the repository in the format "owner/repo".
//
// Returns:
//   - commitList: A slice of GithubCommit structs containing the commit details.
//   - err: An error object if an error occurred during the request or response processing.
//
// The function makes an HTTP GET request to the GitHub API to fetch the commits
// of the specified repository. It sets the necessary headers for authorization
// and content type. If the request is successful, it decodes the JSON response
// into a slice of GithubCommit structs. If any error occurs during the process,
// it returns an appropriate error message.
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
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return commitList, fmt.Errorf("unable to read error response: %w", err)
		}
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

// IsPullRequestUpdate checks if there are any pull requests in the provided list
// that were created after the specified date. It returns true if at least one
// pull request was created after the given date, otherwise it returns false.
//
// Parameters:
//   - pullRequestList: A slice of GithubPullRequest objects to be checked.
//   - date: A time.Time object representing the date to compare against.
//
// Returns:
//   - bool: True if there is at least one pull request created after the specified date, false otherwise.
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

// PullRequestList retrieves a list of pull requests for a given repository from the GitHub API.
//
// Parameters:
//   - userGithubToken: A string containing the GitHub token for authentication.
//   - repo: A string specifying the repository in the format "owner/repo".
//
// Returns:
//   - pullRequestList: A slice of GithubPullRequest structs containing the details of each pull request.
//   - err: An error object if an error occurred during the request or response processing.
//
// The function performs the following steps:
//  1. Creates an HTTP GET request to the GitHub API endpoint for listing pull requests.
//  2. Sets the necessary headers, including the Authorization header with the provided GitHub token.
//  3. Sends the request using the default HTTP client.
//  4. Checks the response status code and reads the response body if the status code is not 200 OK.
//  5. Decodes the JSON response into a slice of GithubPullRequest structs.
//  6. Returns the list of pull requests and any error encountered.
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

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return pullRequestList, fmt.Errorf("unable to read error response: %w", err)
		}
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

// IsWorkflowRunUpdate checks if there is any workflow run in the provided list
// that was created after the specified date. It returns true if at least one
// workflow run was created after the given date, otherwise it returns false.
//
// Parameters:
//   - workflowRunList: A slice of GithubWorkflow objects representing the list
//     of workflow runs to check.
//   - date: A time.Time object representing the date to compare against the
//     creation dates of the workflow runs.
//
// Returns:
//   - bool: True if there is at least one workflow run created after the
//     specified date, otherwise false.
func (service *githubService) IsWorkflowRunUpdate(
	workflowRunList []schemas.GithubWorkflow,
	date time.Time,
) bool {
	for _, workflowRun := range workflowRunList {
		if workflowRun.CreatedAt.After(date) {
			return true
		}
	}

	return false
}

// WorkflowRunList retrieves the list of workflow runs for a given GitHub repository.
// It takes a user's GitHub token and the repository name as parameters and returns
// a list of workflow runs or an error if the request fails.
//
// Parameters:
//   - userGithubToken: A string containing the GitHub token of the user.
//   - repo: A string containing the name of the repository in the format "owner/repo".
//
// Returns:
//   - workflowRunList: A struct containing the list of workflow runs.
//   - err: An error if the request fails or the response cannot be decoded.
func (service *githubService) WorkflowRunList(
	userGithubToken string, repo string,
) (workflowRunList schemas.GithubWorkflowRunsList, err error) {
	ctx := context.Background()

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.github.com/repos/"+repo+"/actions/runs",
		nil,
	)
	if err != nil {
		return workflowRunList, fmt.Errorf("unable to create request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+userGithubToken)
	req.Header.Set("Content-Type", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return workflowRunList, fmt.Errorf("unable to make request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to avoid resource leaks

	if resp.StatusCode != http.StatusOK {
		// Read and log the error response for debugging
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return workflowRunList, fmt.Errorf("unable to read error response: %w", err)
		}
		return workflowRunList, fmt.Errorf(
			"unexpected status code: %d, response: %s",
			resp.StatusCode,
			string(errorBody),
		)
	}

	// Decode the JSON response into the result struct
	err = json.NewDecoder(resp.Body).Decode(&workflowRunList)
	if err != nil {
		return workflowRunList, fmt.Errorf("unable to decode response: %w", err)
	}

	return workflowRunList, nil
}

// Actions functions

// GithubActionUpdateCommitInRepo checks for new commits in a GitHub repository and updates the area storage variable with the latest commit time.
// If a new commit is found, it sends a message to the provided channel.
//
// Parameters:
//   - channel: A channel to send messages about new commits.
//   - option: A JSON raw message containing the options for the GitHub action.
//   - area: The area schema containing user and action information.
//
// The function performs the following steps:
//  1. Retrieves the user's token for the GitHub service.
//  2. Initializes or updates the storage variable with the latest commit time.
//  3. Unmarshals the provided option JSON into a GithubActionOption struct.
//  4. Retrieves the list of commits from the specified repository.
//  5. Checks if there are new commits since the last stored commit time.
//  6. If a new commit is found, updates the storage variable and sends a message to the channel.
//  7. Sleeps for the specified refresh rate before checking again.
func (service *githubService) GithubActionUpdateCommitInRepo(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
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

	databaseStored := schemas.GithubActionOptionStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.GithubActionOptionStorage{
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
		databaseStored = schemas.GithubActionOptionStorage{
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
	optionJSON := schemas.GithubActionOption{}

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

	WaitAction(area)
}

// GithubActionUpdatePullRequestInRepo checks for updates in a GitHub repository's pull requests and sends a notification if there are new updates.
//
// Parameters:
//   - channel: A channel to send the update message.
//   - option: A JSON raw message containing the options for the GitHub action.
//   - area: A schema containing user and action information.
//
// The function performs the following steps:
//  1. Retrieves the user's token for the GitHub service.
//  2. Initializes or updates the storage variable for the action.
//  3. Unmarshals the options for the GitHub action.
//  4. Retrieves the list of pull requests for the specified repository.
//  5. Checks if there are any updates to the pull requests since the last stored time.
//  6. Sends a notification message if there are updates and updates the storage variable.
//  7. Sleeps for the specified refresh rate before the next check.
func (service *githubService) GithubActionUpdatePullRequestInRepo(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
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

	databaseStored := schemas.GithubActionOptionStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.GithubActionOptionStorage{
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
		databaseStored = schemas.GithubActionOptionStorage{
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
	optionJSON := schemas.GithubActionOption{}

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
		response := "new pull request update in " + optionJSON.RepoName + " repository"
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

	WaitAction(area)
}

// GithubActionUpdateWorkflowRunInRepo checks for updates in GitHub workflow runs for a specified repository.
// It retrieves the user's token, unmarshals the storage variable, and updates the storage variable if necessary.
// If a new workflow run is detected, it sends a response message to the provided channel.
//
// Parameters:
//   - channel: A channel to send response messages.
//   - option: A JSON raw message containing the options for the GitHub action.
//   - area: A schema containing user and action information.
//
// The function performs the following steps:
//  1. Retrieves the user's token from the token repository.
//  2. Unmarshals the storage variable from the area schema.
//  3. Initializes the storage variable if it is empty or not properly unmarshalled.
//  4. Unmarshals the option JSON to extract repository information.
//  5. Retrieves the list of workflow runs from the GitHub repository.
//  6. Checks if there are any new workflow runs since the last stored time.
//  7. Updates the storage variable with the current time if a new workflow run is detected.
//  8. Sends a response message to the channel if a new workflow run is detected.
//  9. Sleeps for the specified refresh rate before the next check.
//
// Errors are printed to the console, and the function returns early if any errors occur during the process.
func (service *githubService) GithubActionUpdateWorkflowRunInRepo(
	channel chan string,
	option json.RawMessage,
	area schemas.Area,
) {
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

	databaseStored := schemas.GithubActionOptionStorage{}
	err = json.Unmarshal(area.StorageVariable, &databaseStored)
	if err != nil {
		toto := struct{}{}
		err = json.Unmarshal(area.StorageVariable, &toto)
		if err != nil {
			println("error unmarshalling storage variable: " + err.Error())
			return
		} else {
			println("initializing storage variable")
			databaseStored = schemas.GithubActionOptionStorage{
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
		databaseStored = schemas.GithubActionOptionStorage{
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
	optionJSON := schemas.GithubActionOption{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		println("error unmarshal weather option: " + err.Error())
		time.Sleep(time.Second)
		return
	}

	workflowRunList, err := service.WorkflowRunList(token.Token, optionJSON.RepoName)
	if err != nil {
		fmt.Println(err)
	}

	if service.IsWorkflowRunUpdate(workflowRunList.WorkflowRuns, databaseStored.Time) {
		response := "new workflow run in " + optionJSON.RepoName + " repository"
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

	WaitAction(area)
}

// Reactions functions

// GithubReactionGetLatestCommitInRepo retrieves the latest commit in a specified GitHub repository.
//
// Parameters:
//   - option: A JSON raw message containing the repository name.
//   - area: An Area schema containing user and service information.
//
// Returns:
//   - A string describing the latest commit, including the author's name, commit message, repository name, and commit date.
//   - An error message if any step in the process fails, such as token retrieval, JSON unmarshalling, or commit list retrieval.
func (service *githubService) GithubReactionGetLatestCommitInRepo(
	option json.RawMessage,
	area schemas.Area,
) string {
	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		return "Error finding token:" + err.Error()
	}
	if token.Token == "" {
		return "Error: Token not found"
	}

	// Unmarshal the option
	optionJSON := schemas.GithubActionOption{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		return "error unmarshal weather option: " + err.Error()
	}

	commitList, err := service.CommitList(token.Token, optionJSON.RepoName)
	if err != nil {
		return err.Error()
	}

	if (len(commitList)) == 0 {
		return "No commit found in " + optionJSON.RepoName + " repository"
	} else {
		return commitList[0].Commit.Author.Name + " commit " + commitList[0].Commit.Message + " in " + optionJSON.RepoName + " repository, at " + commitList[0].Commit.Author.Date.String()
	}
}

// GithubReactionGetLatestWorkflowRunInRepo retrieves the latest workflow run in a specified GitHub repository.
//
// Parameters:
//   - option: A JSON raw message containing the repository name.
//   - area: An Area struct containing user and service information.
//
// Returns:
//
//	A string describing the latest workflow run in the specified repository, or an error message if any issues occur.
//
// The function performs the following steps:
//  1. Finds the token of the user for the specified service.
//  2. Unmarshals the option JSON to extract the repository name.
//  3. Retrieves the list of workflow runs for the specified repository.
//  4. Returns the name and creation date of the latest workflow run.
func (service *githubService) GithubReactionGetLatestWorkflowRunInRepo(
	option json.RawMessage,
	area schemas.Area,
) string {
	// Find the token of the user
	token, err := service.tokenRepository.FindByUserIdAndServiceId(
		area.UserId,
		area.Reaction.ServiceId,
	)
	if err != nil {
		return "Error finding token:" + err.Error()
	}
	if token.Token == "" {
		return "Error: Token not found"
	}

	// Unmarshal the option
	optionJSON := schemas.GithubActionOption{}

	err = json.Unmarshal([]byte(option), &optionJSON)
	if err != nil {
		return "error unmarshal weather option: " + err.Error()
	}

	workflowList, err := service.WorkflowRunList(token.Token, optionJSON.RepoName)
	if err != nil {
		return err.Error()
	}

	if workflowList.TotalCount == 0 {
		return "No workflow run found in " + optionJSON.RepoName + " repository"
	} else {
		return workflowList.WorkflowRuns[0].Name + " workflow run in " + optionJSON.RepoName + " repository, at " + workflowList.WorkflowRuns[0].CreatedAt.String()
	}
}

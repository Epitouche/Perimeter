package schemas

import (
	"errors"
	"time"
)

type GithubAction string // GithubAction represents the action to be performed on a Github repository.

// UpdateCommitInRepo is an action to update a commit in a repository.
const (
	UpdateCommitInRepo      GithubAction = "UpdateCommitInRepo"      // UpdateCommitInRepo is an action to update a commit in a repository.
	UpdatePullRequestInRepo GithubAction = "UpdatePullRequestInRepo" // UpdatePullRequestInRepo is an action to update a pull request in a repository.
	UpdateWorkflowRunInRepo GithubAction = "UpdateWorkflowRunInRepo" // UpdateWorkflowRunInRepo is an action to update a workflow run in a repository.
)

type GithubReaction string // GithubReaction represents the reaction to a Github action.

// GetLatestCommitInRepo is a GithubReaction that retrieves the latest commit in a specified repository.
const (
	GetLatestCommitInRepo      GithubReaction = "GetLatestCommitInRepo"      // GetLatestCommitInRepo is a reaction to get the latest commit in a repository.
	GetLatestWorkflowRunInRepo GithubReaction = "GetLatestWorkflowRunInRepo" // GetLatestWorkflowRunInRepo is a reaction to get the latest workflow run in a repository.
)

// GitHubTokenResponse represents the response received from GitHub when
// exchanging a code for an access token. It contains the access token,
// the scope of the token, and the type of the token.
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"` // The access token
	Scope       string `json:"scope"`        // The scope of the token
	TokenType   string `json:"token_type"`   // The type of the token
}

// GithubUserInfo represents the information of a GitHub user.
//
// Fields:
// - Login: The username of the user.
// - Id: The unique identifier of the user. This field is also the primary key in the database.
// - AvatarURL: The URL to the user's avatar.
// - Type: The type of the user.
// - HtmlURL: The URL to the user's profile.
// - Name: The name of the user.
// - Email: The email of the user.
type GithubUserInfo struct {
	Login     string `json:"login"`                        // The username of the user
	Id        uint64 `json:"id"         gorm:"primaryKey"` // The unique identifier of the user
	AvatarURL string `json:"avatar_url"`                   // The URL to the user's avatar
	Type      string `json:"type"`                         // The type of the user
	HtmlURL   string `json:"html_url"`                     // The URL to the user's profile
	Name      string `json:"name"`                         // The name of the user
	Email     string `json:"email"`                        // The email of the user
}

// GithubUserEmail represents the email information of a GitHub user.
// It includes the email address, verification status, primary status,
// and visibility of the email.
type GithubUserEmail struct {
	Email      string `json:"email"`      // The email of the user
	Verified   bool   `json:"verified"`   // Whether the email is verified
	Primary    bool   `json:"primary"`    // Whether the email is the primary email
	Visibility string `json:"visibility"` // The visibility of the email
}

// Errors Messages.
var (
	ErrGithubSecretNotSet             = errors.New("GITHUB_SECRET is not set")               // ErrGithubSecretNotSet is returned when the GITHUB_SECRET environment variable is not set.
	ErrGithubClientIdNotSet           = errors.New("GITHUB_CLIENT_ID is not set")            // ErrGithubClientIdNotSet is returned when the GITHUB_CLIENT_ID environment variable is not set.
	ErrGithubProductionSecretNotSet   = errors.New("GITHUB_PRODUCTION_SECRET is not set")    // ErrGithubProductionSecretNotSet is returned when the GITHUB_PRODUCTION_SECRET environment variable is not set.
	ErrGithubProductionClientIdNotSet = errors.New("GITHUB_PRODUCTION_CLIENT_ID is not set") // ErrGithubProductionClientIdNotSet is returned when the GITHUB_PRODUCTION_CLIENT_ID environment variable is not set.
)

type GithubActionOption struct {
	RepoName string `json:"repo_name"`
}

type GithubActionOptionStorage struct {
	Time time.Time `json:"time"`
}

type GithubActor struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubRepo struct {
	ID               int         `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Owner            GithubActor `json:"owner"`
	Private          bool        `json:"private"`
	HTMLURL          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	ForksURL         string      `json:"forks_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	PullsURL         string      `json:"pulls_url"`
	ReleasesURL      string      `json:"releases_url"`
	SSHURL           string      `json:"ssh_url"`
	StargazersURL    string      `json:"stargazers_url"`
	StatusesURL      string      `json:"statuses_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	TagsURL          string      `json:"tags_url"`
	TeamsURL         string      `json:"teams_url"`
	TreesURL         string      `json:"trees_url"`
	CloneURL         string      `json:"clone_url"`
	MirrorURL        string      `json:"mirror_url"`
	HooksURL         string      `json:"hooks_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Language         interface{} `json:"language"`
	ForksCount       int         `json:"forks_count"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Size             int         `json:"size"`
	DefaultBranch    string      `json:"default_branch"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	IsTemplate       bool        `json:"is_template"`
	Topics           []string    `json:"topics"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	HasDownloads     bool        `json:"has_downloads"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	Visibility       string      `json:"visibility"`
	PushedAt         time.Time   `json:"pushed_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge"`
	TemplateRepository  interface{} `json:"template_repository"`
	TempCloneToken      string      `json:"temp_clone_token"`
	AllowSquashMerge    bool        `json:"allow_squash_merge"`
	AllowAutoMerge      bool        `json:"allow_auto_merge"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge"`
	AllowMergeCommit    bool        `json:"allow_merge_commit"`
	SubscribersCount    int         `json:"subscribers_count"`
	NetworkCount        int         `json:"network_count"`
	License             struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		SpdxID  string `json:"spdx_id"`
		NodeID  string `json:"node_id"`
		HTMLURL string `json:"html_url"`
	} `json:"license"`
	Forks      int `json:"forks"`
	OpenIssues int `json:"open_issues"`
	Watchers   int `json:"watchers"`
}

type GithubCommit struct {
	URL         string `json:"url"`
	Sha         string `json:"sha"`
	NodeID      string `json:"node_id"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Commit      struct {
		URL    string `json:"url"`
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			URL string `json:"url"`
			Sha string `json:"sha"`
		} `json:"tree"`
		CommentCount int `json:"comment_count"`
		Verification struct {
			Verified   bool        `json:"verified"`
			Reason     string      `json:"reason"`
			Signature  interface{} `json:"signature"`
			Payload    interface{} `json:"payload"`
			VerifiedAt interface{} `json:"verified_at"`
		} `json:"verification"`
	} `json:"commit"`
	Author    GithubActor `json:"author"`
	Committer GithubActor `json:"committer"`
	Parents   []struct {
		URL string `json:"url"`
		Sha string `json:"sha"`
	} `json:"parents"`
}

type GithubPullRequest struct {
	URL               string      `json:"url"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	HTMLURL           string      `json:"html_url"`
	DiffURL           string      `json:"diff_url"`
	PatchURL          string      `json:"patch_url"`
	IssueURL          string      `json:"issue_url"`
	CommitsURL        string      `json:"commits_url"`
	ReviewCommentsURL string      `json:"review_comments_url"`
	ReviewCommentURL  string      `json:"review_comment_url"`
	CommentsURL       string      `json:"comments_url"`
	StatusesURL       string      `json:"statuses_url"`
	Number            int         `json:"number"`
	State             string      `json:"state"`
	Locked            bool        `json:"locked"`
	Title             string      `json:"title"`
	User              GithubActor `json:"user"`
	Body              string      `json:"body"`
	Labels            []struct {
		ID          int    `json:"id"`
		NodeID      string `json:"node_id"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		Default     bool   `json:"default"`
	} `json:"labels"`
	Milestone struct {
		URL          string      `json:"url"`
		HTMLURL      string      `json:"html_url"`
		LabelsURL    string      `json:"labels_url"`
		ID           int         `json:"id"`
		NodeID       string      `json:"node_id"`
		Number       int         `json:"number"`
		State        string      `json:"state"`
		Title        string      `json:"title"`
		Description  string      `json:"description"`
		Creator      GithubActor `json:"creator"`
		OpenIssues   int         `json:"open_issues"`
		ClosedIssues int         `json:"closed_issues"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		ClosedAt     time.Time   `json:"closed_at"`
		DueOn        time.Time   `json:"due_on"`
	} `json:"milestone"`
	ActiveLockReason   string        `json:"active_lock_reason"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           time.Time     `json:"closed_at"`
	MergedAt           time.Time     `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	Assignee           GithubActor   `json:"assignee"`
	Assignees          []GithubActor `json:"assignees"`
	RequestedReviewers []GithubActor `json:"requested_reviewers"`
	RequestedTeams     []struct {
		ID                  int         `json:"id"`
		NodeID              string      `json:"node_id"`
		URL                 string      `json:"url"`
		HTMLURL             string      `json:"html_url"`
		Name                string      `json:"name"`
		Slug                string      `json:"slug"`
		Description         string      `json:"description"`
		Privacy             string      `json:"privacy"`
		Permission          string      `json:"permission"`
		NotificationSetting string      `json:"notification_setting"`
		MembersURL          string      `json:"members_url"`
		RepositoriesURL     string      `json:"repositories_url"`
		Parent              interface{} `json:"parent"`
	} `json:"requested_teams"`
	Head struct {
		Label string      `json:"label"`
		Ref   string      `json:"ref"`
		Sha   string      `json:"sha"`
		User  GithubActor `json:"user"`
		Repo  GithubRepo  `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string      `json:"label"`
		Ref   string      `json:"ref"`
		Sha   string      `json:"sha"`
		User  GithubActor `json:"user"`
		Repo  GithubRepo  `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
	AuthorAssociation string      `json:"author_association"`
	AutoMerge         interface{} `json:"auto_merge"`
	Draft             bool        `json:"draft"`
}

type GithubWorkflow struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	NodeID           string        `json:"node_id"`
	CheckSuiteID     int           `json:"check_suite_id"`
	CheckSuiteNodeID string        `json:"check_suite_node_id"`
	HeadBranch       string        `json:"head_branch"`
	HeadSha          string        `json:"head_sha"`
	Path             string        `json:"path"`
	RunNumber        int           `json:"run_number"`
	Event            string        `json:"event"`
	DisplayTitle     string        `json:"display_title"`
	Status           string        `json:"status"`
	Conclusion       interface{}   `json:"conclusion"`
	WorkflowID       int           `json:"workflow_id"`
	URL              string        `json:"url"`
	HTMLURL          string        `json:"html_url"`
	PullRequests     []interface{} `json:"pull_requests"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	Actor            GithubActor   `json:"actor"`
	RunAttempt       int           `json:"run_attempt"`
	RunStartedAt     time.Time     `json:"run_started_at"`
	TriggeringActor  GithubActor   `json:"triggering_actor"`
	JobsURL          string        `json:"jobs_url"`
	LogsURL          string        `json:"logs_url"`
	CheckSuiteURL    string        `json:"check_suite_url"`
	ArtifactsURL     string        `json:"artifacts_url"`
	CancelURL        string        `json:"cancel_url"`
	RerunURL         string        `json:"rerun_url"`
	WorkflowURL      string        `json:"workflow_url"`
	HeadCommit       struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"committer"`
	} `json:"head_commit"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		Private          bool   `json:"private"`
		HTMLURL          string `json:"html_url"`
		Description      string `json:"description"`
		Fork             bool   `json:"fork"`
		URL              string `json:"url"`
		ArchiveURL       string `json:"archive_url"`
		AssigneesURL     string `json:"assignees_url"`
		BlobsURL         string `json:"blobs_url"`
		BranchesURL      string `json:"branches_url"`
		CollaboratorsURL string `json:"collaborators_url"`
		CommentsURL      string `json:"comments_url"`
		CommitsURL       string `json:"commits_url"`
		CompareURL       string `json:"compare_url"`
		ContentsURL      string `json:"contents_url"`
		ContributorsURL  string `json:"contributors_url"`
		DeploymentsURL   string `json:"deployments_url"`
		DownloadsURL     string `json:"downloads_url"`
		EventsURL        string `json:"events_url"`
		ForksURL         string `json:"forks_url"`
		GitCommitsURL    string `json:"git_commits_url"`
		GitRefsURL       string `json:"git_refs_url"`
		GitTagsURL       string `json:"git_tags_url"`
		GitURL           string `json:"git_url"`
		IssueCommentURL  string `json:"issue_comment_url"`
		IssueEventsURL   string `json:"issue_events_url"`
		IssuesURL        string `json:"issues_url"`
		KeysURL          string `json:"keys_url"`
		LabelsURL        string `json:"labels_url"`
		LanguagesURL     string `json:"languages_url"`
		MergesURL        string `json:"merges_url"`
		MilestonesURL    string `json:"milestones_url"`
		NotificationsURL string `json:"notifications_url"`
		PullsURL         string `json:"pulls_url"`
		ReleasesURL      string `json:"releases_url"`
		SSHURL           string `json:"ssh_url"`
		StargazersURL    string `json:"stargazers_url"`
		StatusesURL      string `json:"statuses_url"`
		SubscribersURL   string `json:"subscribers_url"`
		SubscriptionURL  string `json:"subscription_url"`
		TagsURL          string `json:"tags_url"`
		TeamsURL         string `json:"teams_url"`
		TreesURL         string `json:"trees_url"`
		HooksURL         string `json:"hooks_url"`
	} `json:"repository"`
	HeadRepository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL          string      `json:"html_url"`
		Description      interface{} `json:"description"`
		Fork             bool        `json:"fork"`
		URL              string      `json:"url"`
		ForksURL         string      `json:"forks_url"`
		KeysURL          string      `json:"keys_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		TeamsURL         string      `json:"teams_url"`
		HooksURL         string      `json:"hooks_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BranchesURL      string      `json:"branches_url"`
		TagsURL          string      `json:"tags_url"`
		BlobsURL         string      `json:"blobs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		TreesURL         string      `json:"trees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		StargazersURL    string      `json:"stargazers_url"`
		ContributorsURL  string      `json:"contributors_url"`
		SubscribersURL   string      `json:"subscribers_url"`
		SubscriptionURL  string      `json:"subscription_url"`
		CommitsURL       string      `json:"commits_url"`
		GitCommitsURL    string      `json:"git_commits_url"`
		CommentsURL      string      `json:"comments_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		ContentsURL      string      `json:"contents_url"`
		CompareURL       string      `json:"compare_url"`
		MergesURL        string      `json:"merges_url"`
		ArchiveURL       string      `json:"archive_url"`
		DownloadsURL     string      `json:"downloads_url"`
		IssuesURL        string      `json:"issues_url"`
		PullsURL         string      `json:"pulls_url"`
		MilestonesURL    string      `json:"milestones_url"`
		NotificationsURL string      `json:"notifications_url"`
		LabelsURL        string      `json:"labels_url"`
		ReleasesURL      string      `json:"releases_url"`
		DeploymentsURL   string      `json:"deployments_url"`
	} `json:"head_repository"`
}

type GithubWorkflowRunsList struct {
	TotalCount   int              `json:"total_count"`
	WorkflowRuns []GithubWorkflow `json:"workflow_runs"`
}

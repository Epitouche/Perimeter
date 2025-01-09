package schemas

import (
	"errors"
	"time"
)

type DropboxAction string

type DropboxReaction string

const (
	SaveUrl DropboxReaction = "SaveUrl"
)

// DropboxTokenResponse represents the response from Dropbox when a token is requested.
type DropboxTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type DropboxUserInfo struct {
	AccountId   string `json:"account_id"`
	AccountType struct {
		Tag string `json:".tag"`
	} `json:"account_type"`
	Country       string `json:"country"`
	Disabled      bool   `json:"disabled"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	IsPaired      bool   `json:"is_paired"`
	Locale        string `json:"locale"`
	Name          struct {
		AbbreviatedName string `json:"abbreviated_name"`
		DisplayName     string `json:"display_name"`
		FamiliarName    string `json:"familiar_name"`
		GivenName       string `json:"given_name"`
		Surname         string `json:"surname"`
	} `json:"name"`
	ReferralLink string `json:"referral_link"`
	RootInfo     struct {
		Tag             string `json:".tag"`
		HomeNamespaceId string `json:"home_namespace_id"`
		RootNamespaceId string `json:"root_namespace_id"`
	} `json:"root_info"`
}

// Errors Messages.
var (
	ErrDropboxSecretNotSet   = errors.New("DROPBOX_SECRET is not set")
	ErrDropboxClientIdNotSet = errors.New("DROPBOX_CLIENT_ID is not set")
)

type DropboxMobileTokenRequest struct {
	Token string `json:"token"`
}

type DropboxFile struct {
	Created     time.Time `json:"created"`
	Destination string    `json:"destination"`
	FileCount   int       `json:"file_count"`
	ID          string    `json:"id"`
	IsOpen      bool      `json:"is_open"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
}

type DropboxListFileRequestsV2Result struct {
	Cursor       string        `json:"cursor"`
	FileRequests []DropboxFile `json:"file_requests"`
	HasMore      bool          `json:"has_more"`
}

type DropboxEntry struct {
	Tag            string    `json:".tag"`
	ClientModified time.Time `json:"client_modified,omitempty"`
	ContentHash    string    `json:"content_hash,omitempty"`
	FileLockInfo   struct {
		Created        time.Time `json:"created"`
		IsLockholder   bool      `json:"is_lockholder"`
		LockholderName string    `json:"lockholder_name"`
	} `json:"file_lock_info,omitempty"`
	HasExplicitSharedMembers bool   `json:"has_explicit_shared_members,omitempty"`
	ID                       string `json:"id"`
	IsDownloadable           bool   `json:"is_downloadable,omitempty"`
	Name                     string `json:"name"`
	PathDisplay              string `json:"path_display"`
	PathLower                string `json:"path_lower"`
	PropertyGroups           []struct {
		Fields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
		TemplateID string `json:"template_id"`
	} `json:"property_groups"`
	Rev            string    `json:"rev,omitempty"`
	ServerModified time.Time `json:"server_modified,omitempty"`
	SharingInfo    struct {
		ModifiedBy           string `json:"modified_by"`
		ParentSharedFolderID string `json:"parent_shared_folder_id"`
		ReadOnly             bool   `json:"read_only"`
	} `json:"sharing_info"`
	Size int `json:"size,omitempty"`
}

type DropboxListFolderResult struct {
	Cursor  string         `json:"cursor"`
	Entries []DropboxEntry `json:"entries"`
	HasMore bool           `json:"has_more"`
}

type DropboxCountFileRequestsResult struct {
	FileRequestCount uint64 `json:"file_request_count"`
}

type DropboxSaveUrlReactionOption struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

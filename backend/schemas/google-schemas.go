package schemas

import (
	"errors"
	"time"
)

type GoogleAction string // The action type for Google.

const (
	ReceiveGoogleMail GoogleAction = "ReceiveGoogleMail" // ReceiveGoogleMail is the action to receive emails from Google.
)

type GoogleReaction string // The reaction type for Google.

const (
	SendMail GoogleReaction = "SendMail"
)

// GoogleTokenResponse represents the response from Google's OAuth 2.0 token endpoint.
// It contains the access token, expiration time, scope, token type, and refresh token.
type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`  // The access token
	ExpiresIn    uint64 `json:"expires_in"`    // The expiration time of the token in seconds
	Scope        string `json:"scope"`         // The scope of the token
	TokenType    string `json:"token_type"`    // The type of the token
	RefreshToken string `json:"refresh_token"` // The refresh token
}

// GmailProfile represents a user's Gmail profile information.
// It includes the user's email address, the total number of messages,
// the total number of threads, and the history ID.
type GmailProfile struct {
	EmailAddress  string `json:"emailAddress"`  // The email address of the user
	MessagesTotal int    `json:"messagesTotal"` // The total number of messages
	ThreadsTotal  int    `json:"threadsTotal"`  // The total number of threads
	HistoryId     string `json:"historyId"`     // The history ID
}

// GmailUserInfo represents the information of a Gmail user.
// It contains the login email and the email address of the user.
type GmailUserInfo struct {
	Login string `json:"login"` // The login email of the user
	Email string `json:"email"` // The email of the user
}

type GoogleProfile struct {
	ResourceName string `json:"resourceName"`
	Etag         string `json:"etag"`
	Names        []struct {
		Metadata struct {
			Primary bool `json:"primary"`
			Source  struct {
				Type string `json:"type"`
				Id   string `json:"id"`
			} `json:"source"`
			SourcePrimary bool `json:"sourcePrimary"`
		} `json:"metadata"`
		DisplayName          string `json:"displayName"`
		GivenName            string `json:"givenName"`
		DisplayNameLastFirst string `json:"displayNameLastFirst"`
		UnstructuredName     string `json:"unstructuredName"`
	} `json:"names"`
}

type GmailReactionSendMailOption struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// Errors Messages.
var (
	ErrGoogleSecretNotSet   = errors.New("GOOGLE_SECRET is not set")
	ErrGoogleClientIdNotSet = errors.New("GOOGLE_CLIENT_ID is not set")
)

type GoogleVariableReceiveMail struct {
	Time time.Time `json:"time"`
}

type GmailMessage struct {
	Id     string `json:"id"`
	Thread string `json:"threadId"`
}

type GmailEmailResponse struct {
	Messages []GmailMessage `json:"messages"`
}

type EmailDetails struct {
	Date    string `json:"date"`
	From    string `json:"from"`
	Subject string `json:"body"`
}

type GmailMessageResponse struct {
	Payload struct {
		Headers []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
		Parts []struct {
			MimeType string `json:"mimeType"`
			Body     struct {
				Data string `json:"data"`
			} `json:"body"`
		} `json:"parts"`
	} `json:"payload"`
}

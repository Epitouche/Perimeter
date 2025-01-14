package schemas

import (
	"errors"
	"time"
)

type GoogleAction string

const (
	ReceiveGoogleMail GoogleAction = "ReceiveGoogleMail"
)

type GoogleReaction string

const (
	SendMail GoogleReaction = "SendMail"
)

// GoogleTokenResponse represents the response from Gmail when a token is requested.
type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type GmailProfile struct {
	EmailAddress  string `json:"emailAddress"`
	MessagesTotal int    `json:"messagesTotal"`
	ThreadsTotal  int    `json:"threadsTotal"`
	HistoryId     string `json:"historyId"`
}

type GmailUserInfo struct {
	Login string `json:"login"`
	Email string `json:"email"`
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

package schemas

type GmailAction string

const ()

type GmailReaction string

const (
	SendMail GmailReaction = "SendMail"
)

// GmailTokenResponse represents the response from Gmail when a token is requested.
type GmailTokenResponse struct {
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

type GmailProfileNamesMetadataSource struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type GmailProfileNamesMetadata struct {
	Primary       bool                            `json:"primary"`
	Source        GmailProfileNamesMetadataSource `json:"source"`
	SourcePrimary bool                            `json:"sourcePrimary"`
}

type GoogleProfileNames struct {
	Metadata             GmailProfileNamesMetadata `json:"metadata"`
	DisplayName          string                    `json:"displayName"`
	GivenName            string                    `json:"givenName"`
	DisplayNameLastFirst string                    `json:"displayNameLastFirst"`
	UnstructuredName     string                    `json:"unstructuredName"`
}

type GoogleProfile struct {
	ResourceName string               `json:"resourceName"`
	Etag         string               `json:"etag"`
	Names        []GoogleProfileNames `json:"names"`
}

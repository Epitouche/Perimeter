package schemas

import "errors"

type DiscordAction string

type DiscordReaction string

// error messages
var (
	ErrDiscordClientIdNotSet = errors.New("DISCORD_CLIENT_ID is not set")
	ErrDiscordSecretNotSet   = errors.New("DISCORD_SECRET is not set")
)
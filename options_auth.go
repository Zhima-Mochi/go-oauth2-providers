package oauth2providers

type AuthOption func(*authConfig)

var (
	defaultAuthConfig = &authConfig{}
)

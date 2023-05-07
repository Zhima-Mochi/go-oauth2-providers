package oauth2providers

type AuthConfig interface {
}

type authConfig struct {
}

func NewAuthConfig(opts ...AuthOption) AuthConfig {
	ac := defaultAuthConfig
	for _, opt := range opts {
		opt(ac)
	}

	return ac
}

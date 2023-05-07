package oauth2providers

type ProviderOption func(*providerConfig)

func WithProviderScopes(scopes []string) ProviderOption {
	return func(cfg *providerConfig) {
		cfg.scopes = scopes
	}
}

func WithProviderClientID(clientID string) ProviderOption {
	return func(cfg *providerConfig) {
		cfg.clientID = clientID
	}
}

func WithProviderClientSecret(clientSecret string) ProviderOption {
	return func(cfg *providerConfig) {
		cfg.clientSecret = clientSecret
	}
}

func WithProviderRedirectURL(redirectURL string) ProviderOption {
	return func(cfg *providerConfig) {
		cfg.redirectURL = redirectURL
	}
}

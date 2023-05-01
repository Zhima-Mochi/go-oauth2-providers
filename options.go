package oauth2providers

import (
	"golang.org/x/oauth2"
)

type AuthOption func(*authOptions)

type authOptions struct {
	oauth2.Config
}

// WithScopes sets the Scopes field of the oauth2 Config
func WithScopes(scopes []string) AuthOption {
	return func(os *authOptions) {
		os.Scopes = scopes
	}
}

// WithClientID sets the ClientID field of the oauth2 Config
func WithClientID(clientID string) AuthOption {
	return func(os *authOptions) {
		os.ClientID = clientID
	}
}

// WithClientSecret sets the ClientSecret field of the oauth2 Config
func WithClientSecret(clientSecret string) AuthOption {
	return func(os *authOptions) {
		os.ClientSecret = clientSecret
	}
}

// WithRedirectURL sets the RedirectURL field of the oauth2 Config
func WithRedirectURL(redirectURL string) AuthOption {
	return func(os *authOptions) {
		os.RedirectURL = redirectURL
	}
}

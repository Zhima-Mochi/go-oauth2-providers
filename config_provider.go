package oauth2providers

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

var (

	// ErrMissingClientID is returned when the client ID is missing.
	ErrMissingClientID = errors.New("missing client ID")

	// ErrMissingClientSecret is returned when the client secret is missing.
	ErrMissingClientSecret = errors.New("missing client secret")

	// ErrMissingRedirectURL is returned when the redirect URL is missing.
	ErrMissingRedirectURL = errors.New("missing redirect URL")
)

type ProviderConfig interface {
	AuthCodeURL(state string, opts ...AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error)
	TokenSource(ctx context.Context, token *Token) oauth2.TokenSource
	Client(ctx context.Context, token *Token) *http.Client
	addScopes(scopes ...string)
	setAuthURL(authURL string)
	setTokenURL(tokenURL string)
}

type providerConfig struct {
	authURL      string
	tokenURL     string
	scopes       []string
	clientID     string
	clientSecret string
	redirectURL  string
	oauth2.Config
}

func validateProviderConfig(pc *providerConfig) error {
	if pc.clientID == "" {
		return ErrMissingClientID
	}
	if pc.clientSecret == "" {
		return ErrMissingClientSecret
	}
	if pc.redirectURL == "" {
		return ErrMissingRedirectURL
	}
	return nil
}

func NewProviderConfig(opts ...ProviderOption) (ProviderConfig, error) {
	pc := &providerConfig{}
	for _, opt := range opts {
		opt(pc)
	}

	if err := validateProviderConfig(pc); err != nil {
		return nil, err
	}

	pc.Config = oauth2.Config{
		ClientID:     pc.clientID,
		ClientSecret: pc.clientSecret,
		RedirectURL:  pc.redirectURL,
		Scopes:       pc.scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  pc.authURL,
			TokenURL: pc.tokenURL,
		},
	}
	return pc, nil
}

func (pc *providerConfig) AuthCodeURL(state string, opts ...AuthCodeOption) string {
	return pc.Config.AuthCodeURL(state, convertAuthCodeOptions(opts)...)
}

func (pc *providerConfig) Exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error) {
	return pc.Config.Exchange(ctx, code, convertAuthCodeOptions(opts)...)
}

func (pc *providerConfig) TokenSource(ctx context.Context, token *Token) oauth2.TokenSource {
	return pc.Config.TokenSource(ctx, token)
}

func (pc *providerConfig) Client(ctx context.Context, token *Token) *http.Client {
	return pc.Config.Client(ctx, token)
}

func (pc *providerConfig) addScopes(scopes ...string) {
	hasSeen := map[string]bool{}
	for _, scope := range pc.scopes {
		hasSeen[scope] = true
	}

	for _, scope := range scopes {
		if _, ok := hasSeen[scope]; !ok {
			pc.scopes = append(pc.scopes, scope)
			hasSeen[scope] = true
		}
	}

	pc.Config.Scopes = pc.scopes
}

func (pc *providerConfig) setAuthURL(authURL string) {
	pc.authURL = authURL

	pc.Config.Endpoint.AuthURL = authURL
}

func (pc *providerConfig) setTokenURL(tokenURL string) {
	pc.tokenURL = tokenURL

	pc.Config.Endpoint.TokenURL = tokenURL
}

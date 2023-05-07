package oauth2providers

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type ProviderConfig interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource
	Client(ctx context.Context, token *oauth2.Token) *http.Client
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

func NewProviderConfig(opts ...ProviderOption) ProviderConfig {
	pc := &providerConfig{}
	for _, opt := range opts {
		opt(pc)
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
	return pc
}

func (pc *providerConfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return pc.Config.AuthCodeURL(state, opts...)
}

func (pc *providerConfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return pc.Config.Exchange(ctx, code, opts...)
}

func (pc *providerConfig) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	return pc.Config.TokenSource(ctx, token)
}

func (pc *providerConfig) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return pc.Config.Client(ctx, token)
}

func (pc *providerConfig) addScopes(scopes ...string) {
	for _, scope := range scopes {
		flag := false
		for _, s := range pc.scopes {
			if s == scope {
				flag = true
				break
			}
		}
		if !flag {
			pc.scopes = append(pc.scopes, scope)
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

package provider

import (
	"context"

	"golang.org/x/oauth2"
)

type ProviderType string

const (
	GoogleOAuth2ProviderType   ProviderType = "google"
	FacebookOAuth2ProviderType ProviderType = "facebook"
	LineOAuth2ProviderType     ProviderType = "line"
	// GithubOAuth2ProviderType   ProviderType = "github"
)

type Provider struct {
	providerType ProviderType
	config       *oauth2.Config
}

func NewProvider(providerType ProviderType, clientID, clientSecret, redirectURL string) *Provider {
	var provider *Provider
	switch providerType {
	case GoogleOAuth2ProviderType:
		provider = newGoogle(
			&oauth2.Config{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RedirectURL:  redirectURL,
			},
		)
	case FacebookOAuth2ProviderType:
		provider = newFacebook(&oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
		})
	case LineOAuth2ProviderType:
		provider = newLine(&oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
		})
	default:
		panic("Invalid OAuth2 provider type")
	}
	provider.providerType = providerType
	return provider
}

func (p *Provider) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.config.AuthCodeURL(state, opts...)
}

func (p *Provider) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.config.Exchange(ctx, code, opts...)
}

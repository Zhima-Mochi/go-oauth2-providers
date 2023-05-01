package oauth2providers

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
	authOptions  *authOptions
}

type UserInfo struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func NewProvider(providerType ProviderType, authOptions *authOptions) *Provider {
	var provider *Provider
	switch providerType {
	case GoogleOAuth2ProviderType:
		provider = newGoogle(
			authOptions,
		)
	case FacebookOAuth2ProviderType:
		provider = newFacebook(
			authOptions,
		)
	case LineOAuth2ProviderType:
		provider = newLine(
			authOptions,
		)
	default:
		panic("Invalid OAuth2 provider type")
	}
	provider.providerType = providerType
	return provider
}

func (p *Provider) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.authOptions.AuthCodeURL(state, opts...)
}

func (p *Provider) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.authOptions.Exchange(ctx, code, opts...)
}

func (p *Provider) GetUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error) {
	switch p.providerType {
	case GoogleOAuth2ProviderType:
		return getGoogleUserInfo(ctx, p.authOptions, token)
	case FacebookOAuth2ProviderType:
		return getFacebookUserInfo(ctx, p.authOptions, token)
	case LineOAuth2ProviderType:
		return getLineUserInfo(ctx, p.authOptions, token)
	default:
		panic("Invalid OAuth2 provider type")
	}
}

// refresh
func (p *Provider) RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return p.authOptions.TokenSource(ctx, token).Token()
}

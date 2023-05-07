package oauth2providers

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type ProviderType int

const (
	GoogleOAuth2ProviderType ProviderType = iota
	FacebookOAuth2ProviderType
	LineOAuth2ProviderType
	// GithubOAuth2ProviderType
)

type OAuth2Provider interface {
	authCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	getUserInfo(ctx context.Context, token *oauth2.Token) (UserInfo, error)
	refreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error)
}

func (t ProviderType) String() string {
	switch t {
	case GoogleOAuth2ProviderType:
		return "Google"
	case FacebookOAuth2ProviderType:
		return "Facebook"
	case LineOAuth2ProviderType:
		return "Line"
	default:
		return ""
	}
}

func NewOAuth2Provider(providerType ProviderType, providerConfig ProviderConfig) (OAuth2Provider, error) {
	var provider OAuth2Provider
	switch providerType {
	case GoogleOAuth2ProviderType:
		provider = newGoogleProvider(providerConfig)
	case FacebookOAuth2ProviderType:
		provider = newFacebookProvider(providerConfig)
	case LineOAuth2ProviderType:
		provider = newLineProvider(providerConfig)
	default:
		return nil, fmt.Errorf("unknown OAuth2 provider type: %s", providerType)
	}
	return provider, nil
}

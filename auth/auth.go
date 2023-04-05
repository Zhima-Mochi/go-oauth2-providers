package auth

import (
	"context"

	provider "github.com/Zhima-Mochi/go-auth-providers/providers"
	"golang.org/x/oauth2"
)

type Oauth2 struct {
	provider *provider.Provider
}

func NewOauth2(providerType provider.ProviderType, clientID, clientSecret, redirectURL string) *Oauth2 {
	return &Oauth2{
		provider: provider.NewProvider(providerType, clientID, clientSecret, redirectURL),
	}
}

func (h *Oauth2) Login(ctx context.Context, options ...oauth2.AuthCodeOption) string {
	url := h.provider.AuthCodeURL("state", options...)
	return url
}

func (h *Oauth2) Callback(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := h.provider.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (h *Oauth2) Refresh(ctx context.Context, refreshToken *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := h.provider.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

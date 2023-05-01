package oauth2providers

import (
	"context"

	"golang.org/x/oauth2"
)

type Auth struct {
	*Provider
}

func NewAuth(providerType ProviderType, options ...AuthOption) *Auth {
	authOptions := &authOptions{}
	for _, option := range options {
		option(authOptions)
	}
	return &Auth{
		Provider: NewProvider(providerType, authOptions),
	}
}

func (a *Auth) Login(ctx context.Context, options ...oauth2.AuthCodeOption) string {
	url := a.AuthCodeURL("state", options...)
	return url
}

func (a *Auth) Callback(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := a.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *Auth) Refresh(ctx context.Context, refreshToken *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := a.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

package oauth2providers

import (
	"context"

	"golang.org/x/oauth2"
)

type Auth interface {
	Login(ctx context.Context, options ...oauth2.AuthCodeOption) string
	Callback(ctx context.Context, code string) (*oauth2.Token, error)
	Refresh(ctx context.Context, refreshToken *oauth2.Token) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (UserInfo, error)
}

type auth struct {
	provider OAuth2Provider
}

func NewAuth(provider OAuth2Provider) Auth {
	return &auth{
		provider: provider,
	}
}

func (a *auth) Login(ctx context.Context, options ...oauth2.AuthCodeOption) string {
	url := a.provider.authCodeURL("state", options...)
	return url
}

func (a *auth) Callback(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := a.provider.exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *auth) Refresh(ctx context.Context, refreshToken *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := a.provider.refreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func (a *auth) GetUserInfo(ctx context.Context, token *oauth2.Token) (UserInfo, error) {
	userInfo, err := a.provider.getUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

package oauth2providers

import (
	"context"

	"github.com/google/uuid"
)

var (
	generateStateToken = func() string {
		return uuid.New().String()
	}
)

type Auth interface {
	Login(ctx context.Context, options ...AuthCodeOption) string
	Callback(ctx context.Context, code string) (Token, error)
	Refresh(ctx context.Context, refreshToken Token) (Token, error)
	GetUserInfo(ctx context.Context, token Token) (UserInfo, error)
}

type auth struct {
	AuthConfig
	provider OAuth2Provider
}

func NewAuth(provider OAuth2Provider, authConfig AuthConfig) Auth {
	a := &auth{
		AuthConfig: authConfig,
		provider:   provider,
	}
	return a
}

func (a *auth) Login(ctx context.Context, options ...AuthCodeOption) string {
	state := generateStateToken()
	url := a.provider.authCodeURL(state, options...)
	return url
}

func (a *auth) Callback(ctx context.Context, code string) (Token, error) {
	token, err := a.provider.exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *auth) Refresh(ctx context.Context, refreshToken Token) (Token, error) {
	newToken, err := a.provider.refreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func (a *auth) GetUserInfo(ctx context.Context, token Token) (UserInfo, error) {
	userInfo, err := a.provider.getUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

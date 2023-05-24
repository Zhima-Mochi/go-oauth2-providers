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
	GetOAuth2AuthCodeURL(ctx context.Context, options ...AuthCodeOption) (url string)
	ExchangeOAuth2AuthCode(ctx context.Context, code string) (token *Token, err error)
	RefreshOAuth2Token(ctx context.Context, refreshToken *Token) (token *Token, err error)
	GetOAuth2UserInfo(ctx context.Context, token *Token) (userInfo UserInfo, err error)
}

type auth struct {
	AuthConfig
	provider OAuth2Provider
}

func NewOAuth2Auth(provider OAuth2Provider, authConfig AuthConfig) Auth {
	a := &auth{
		AuthConfig: authConfig,
		provider:   provider,
	}
	return a
}

func (a *auth) GetOAuth2AuthCodeURL(ctx context.Context, options ...AuthCodeOption) (url string) {
	state := generateStateToken()
	url = a.provider.authCodeURL(state, options...)
	return url
}

func (a *auth) ExchangeOAuth2AuthCode(ctx context.Context, code string) (token *Token, err error) {
	token, err = a.provider.exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *auth) RefreshOAuth2Token(ctx context.Context, refreshToken *Token) (token *Token, err error) {
	newToken, err := a.provider.refreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	token = newToken
	return token, nil
}

func (a *auth) GetOAuth2UserInfo(ctx context.Context, token *Token) (userInfo UserInfo, err error) {
	userInfo, err = a.provider.getUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

package oauth2providers

import (
	"context"
)

type lineProvider struct {
	ProviderConfig
}

func (p *lineProvider) authCodeURL(state string, opts ...AuthCodeOption) string {
	return p.AuthCodeURL(state, opts...)
}

func (p *lineProvider) exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error) {
	return p.Exchange(ctx, code, opts...)
}

func (p *lineProvider) refreshToken(ctx context.Context, token *Token) (*Token, error) {
	return p.TokenSource(ctx, token).Token()
}

func newLineProvider(config ProviderConfig) *lineProvider {
	// line endpoint
	config.setAuthURL("https://access.line.me/oauth2/v2.1/authorize")
	config.setTokenURL("https://api.line.me/oauth2/v2.1/token")
	config.addScopes("profile", "openid", "email")

	return &lineProvider{
		ProviderConfig: config,
	}
}

func (p *lineProvider) getUserInfo(ctx context.Context, token *Token) (UserInfo, error) {
	client := p.Client(ctx, token)
	resp, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user, err := parseJSONFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := NewUserInfo()
	if id, ok := user["userId"]; ok {
		userInfo.setID(id.(string))
	}
	if name, ok := user["displayName"]; ok {
		userInfo.setName(name.(string))
	}
	if picture, ok := user["pictureUrl"]; ok {
		userInfo.setPictureURL(picture.(string))
	}
	if email, ok := user["email"]; ok {
		userInfo.setEmail(email.(string))
	}

	return userInfo, nil
}

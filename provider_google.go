package oauth2providers

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleProvider struct {
	ProviderConfig ProviderConfig
}

func (p *googleProvider) authCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.ProviderConfig.AuthCodeURL(state, opts...)
}

func (p *googleProvider) exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.ProviderConfig.Exchange(ctx, code, opts...)
}

func (p *googleProvider) refreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return p.ProviderConfig.TokenSource(ctx, token).Token()
}

func newGoogleProvider(config ProviderConfig) *googleProvider {
	// google endpoint
	config.setAuthURL(google.Endpoint.AuthURL)
	config.setTokenURL(google.Endpoint.TokenURL)
	config.addScopes("https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile")
	return &googleProvider{
		ProviderConfig: config,
	}
}

func (p *googleProvider) getUserInfo(ctx context.Context, token *oauth2.Token) (UserInfo, error) {
	client := p.ProviderConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user, err := parseJSONFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := NewUserInfo()
	if id, ok := user["sub"]; ok {
		userInfo.setID(id.(string))
	}
	if name, ok := user["name"]; ok {
		userInfo.setName(name.(string))
	}
	if email, ok := user["email"]; ok {
		userInfo.setEmail(email.(string))
	}
	if picture, ok := user["picture"]; ok {
		userInfo.setPictureURL(picture.(string))
	}
	return userInfo, nil
}

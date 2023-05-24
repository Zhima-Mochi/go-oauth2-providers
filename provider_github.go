package oauth2providers

import (
	"context"
	"strconv"

	"golang.org/x/oauth2/github"
)

type githubProvider struct {
	ProviderConfig
}

func (p *githubProvider) authCodeURL(state string, opts ...AuthCodeOption) string {
	return p.AuthCodeURL(state, opts...)
}

func (p *githubProvider) exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error) {
	return p.Exchange(ctx, code, opts...)
}

func (p *githubProvider) refreshToken(ctx context.Context, token *Token) (*Token, error) {
	return p.TokenSource(ctx, token).Token()
}

func newGithubProvider(config ProviderConfig) *githubProvider {
	// github endpoint
	config.setAuthURL(github.Endpoint.AuthURL)
	config.setTokenURL(github.Endpoint.TokenURL)
	config.addScopes("read:user", "user:email")
	return &githubProvider{
		ProviderConfig: config,
	}
}

func (p *githubProvider) getUserInfo(ctx context.Context, token *Token) (UserInfo, error) {
	client := p.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user, err := parseJSONFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	userInfo := NewUserInfo()
	if id, ok := user["id"]; ok {
		userInfo.setID(strconv.FormatFloat(id.(float64), 'f', 0, 64))
	}
	if login, ok := user["login"]; ok {
		userInfo.setName(login.(string))
	}
	if email, ok := user["email"]; ok && email != nil {
		userInfo.setEmail(email.(string))
	}
	if picture, ok := user["avatar_url"]; ok {
		userInfo.setPictureURL(picture.(string))
	}
	return userInfo, nil
}

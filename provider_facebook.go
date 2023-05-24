package oauth2providers

import (
	"context"

	"golang.org/x/oauth2/facebook"
)

type facebookProvider struct {
	ProviderConfig
}

func (p *facebookProvider) authCodeURL(state string, opts ...AuthCodeOption) string {
	return p.AuthCodeURL(state, opts...)
}

func (p *facebookProvider) exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error) {
	return p.Exchange(ctx, code, opts...)
}

func (p *facebookProvider) refreshToken(ctx context.Context, token *Token) (*Token, error) {
	return p.TokenSource(ctx, token).Token()
}

func (p *facebookProvider) getUserInfo(ctx context.Context, token *Token) (UserInfo, error) {
	client := p.Client(ctx, token)
	resp, err := client.Get("https://graph.facebook.com/v2.12/me?fields=id,email,name,picture")
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
		userInfo.setID(id.(string))
	}
	if name, ok := user["name"]; ok {
		userInfo.setName(name.(string))
	}
	if email, ok := user["email"]; ok {
		userInfo.setEmail(email.(string))
	}
	if picture, ok := user["picture"]; ok {
		if pictureMap, ok := picture.(map[string]interface{}); ok {
			if data, ok := pictureMap["data"].(map[string]interface{}); ok {
				if url, ok := data["url"].(string); ok {
					userInfo.setPictureURL(url)
				}
			}
		}
	}
	return userInfo, nil
}

func newFacebookProvider(config ProviderConfig) *facebookProvider {
	// facebook endpoint
	config.setAuthURL(facebook.Endpoint.AuthURL)
	config.setTokenURL(facebook.Endpoint.TokenURL)
	config.addScopes("email", "public_profile")
	return &facebookProvider{
		ProviderConfig: config,
	}
}

package oauth2providers

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type facebookProvider struct {
	ProviderConfig ProviderConfig
}

func (p *facebookProvider) authCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.ProviderConfig.AuthCodeURL(state, opts...)
}

func (p *facebookProvider) exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.ProviderConfig.Exchange(ctx, code, opts...)
}

func (p *facebookProvider) refreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return p.ProviderConfig.TokenSource(ctx, token).Token()
}

func (p *facebookProvider) getUserInfo(ctx context.Context, token *oauth2.Token) (UserInfo, error) {
	client := p.ProviderConfig.Client(ctx, token)
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

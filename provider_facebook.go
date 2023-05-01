package oauth2providers

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func newFacebook(options *authOptions) *Provider {
	options.Endpoint = facebook.Endpoint
	if options.Scopes == nil {
		options.Scopes = []string{
			"email",
			"public_profile",
		}
	}
	return &Provider{
		authOptions: options,
	}
}

func getFacebookUserInfo(ctx context.Context, options *authOptions, accessToken *oauth2.Token) (*UserInfo, error) {
	client := options.Client(ctx, accessToken)
	resp, err := client.Get("https://graph.facebook.com/v2.12/me?fields=id,email,name,picture")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		Email:         userInfo.Email,
		EmailVerified: true,
		Name:          userInfo.Name,
		Picture:       userInfo.Picture,
		Locale:        "",
	}, nil
}

package provider

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func newFacebook(config *oauth2.Config) *Provider {
	if config == nil {
		config = &oauth2.Config{}
	}
	config.Endpoint = facebook.Endpoint
	config.Scopes = []string{
		"email",
		"public_profile",
	}
	return &Provider{
		config: config,
	}
}

func getFacebookUserInfo(ctx context.Context, config *oauth2.Config, accessToken *oauth2.Token) (*UserInfo, error) {
	client := config.Client(ctx, accessToken)
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

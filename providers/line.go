package provider

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
)

func newLine(config *oauth2.Config) *Provider {
	if config == nil {
		config = &oauth2.Config{}
	}
	// line endpoint
	config.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
		TokenURL: "https://api.line.me/oauth2/v2.1/token",
	}
	config.Scopes = []string{
		"profile",
		"openid",
		"email",
	}

	return &Provider{
		config: config,
	}
}

func getLineUserInfo(ctx context.Context, config *oauth2.Config, accessToken *oauth2.Token) (*UserInfo, error) {
	client := config.Client(ctx, accessToken)
	resp, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID      string `json:"userId"`
		Email   string `json:"email"`
		Name    string `json:"displayName"`
		Picture string `json:"pictureUrl"`
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

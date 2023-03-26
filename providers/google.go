package provider

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func newGoogle(config *oauth2.Config) *Provider {
	if config == nil {
		config = &oauth2.Config{}
	}
	config.Endpoint = google.Endpoint
	config.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	return &Provider{
		config: config,
	}
}

func getGoogleUserInfo(ctx context.Context, config *oauth2.Config, accessToken *oauth2.Token) (*UserInfo, error) {
	return nil, nil
}

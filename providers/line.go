package provider

import (
	"context"

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
	return nil, nil
}

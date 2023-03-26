package provider

import (
	"context"

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
	return nil, nil
}

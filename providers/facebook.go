package provider

import (
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

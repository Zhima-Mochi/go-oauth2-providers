package oauth2providers

import "golang.org/x/oauth2"

type AuthCodeOption struct {
	Key string
	Val string
}

func convertAuthCodeOptions(opts []AuthCodeOption) []oauth2.AuthCodeOption {
	oauth2Opts := make([]oauth2.AuthCodeOption, len(opts))
	for i, opt := range opts {
		oauth2Opts[i] = oauth2.SetAuthURLParam(opt.Key, opt.Val)
	}
	return oauth2Opts
}

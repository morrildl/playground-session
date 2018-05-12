/* Copyright © Playground Global, LLC. All rights reserved. */

package session

type OAuthConfig struct {
	Issuer           string
	ClientID         string
	ClientSecret     string
	RedirectURL      string
	RedirectPath     string
	Scopes           []string
	AuthURL          string
	TokenExchangeURL string
	JWTPubKeyURL     string
	ValidEmailRegex  string
}

type ConfigType struct {
	SessionCookieKey string
	OAuth            OAuthConfig
}

var Config ConfigType = ConfigType{
	"X-Playground-Session",
	OAuthConfig{
		"accounts.example.com",
		"client_id_aka_audience",
		"client_secret",
		"http://localhost:9000/oauth",
		"/oauth",
		[]string{"openid", "email"},
		"https://oauth.example.com/auth",
		"https://oauth.example.com/token",
		"https://oauth.example.com/keys",
		".*",
	},
}

package oauth

import (
	"envoy-golang-filter-hub/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GitHub *oauth2.Config

func Init() {
	GitHub = &oauth2.Config{
		ClientID:     config.Get().OAuth.ClientID,
		ClientSecret: config.Get().OAuth.ClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  config.Get().OAuth.RedirectURL,
		Scopes:       []string{config.Get().OAuth.Scope},
	}
}

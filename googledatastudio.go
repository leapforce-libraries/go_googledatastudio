package googledatastudio

import (
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"

	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName         string = "GoogleDataStudio"
	apiURL          string = "https://datastudio.googleapis.com/v1"
	authURL         string = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenURL        string = "https://oauth2.googleapis.com/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
)

// GoogleDataStudio stores GoogleDataStudio configuration
//
type GoogleDataStudio struct {
	oAuth2 *oauth2.OAuth2
}

// methods
//
func NewGoogleDataStudio(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery, isLive bool) (*GoogleDataStudio, error) {
	gd := GoogleDataStudio{}
	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Scope:           scope,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
	}
	gd.oAuth2 = oauth2.NewOAuth(config, bigQuery, isLive)
	return &gd, nil
}

func (gc *GoogleDataStudio) InitToken() error {
	return gc.oAuth2.InitToken()
}

func (gd *GoogleDataStudio) Get(url string, model interface{}) (*http.Response, error) {
	res, err := gd.oAuth2.Get(url, model)

	if err != nil {
		return nil, err
	}

	return res, nil
}

package googledatastudio

import (
	"bytes"
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
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
func NewGoogleDataStudio(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleDataStudio {
	gd := GoogleDataStudio{}
	maxRetries := uint(3)
	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Scope:           scope,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
		MaxRetries:      &maxRetries,
	}
	gd.oAuth2 = oauth2.NewOAuth(config, bigQuery)
	return &gd
}

func (gc *GoogleDataStudio) InitToken() *errortools.Error {
	return gc.oAuth2.InitToken()
}

func (gd *GoogleDataStudio) get(url string, model interface{}) (*http.Request, *http.Response, *errortools.Error) {
	err := google.ErrorResponse{}
	request, response, e := gd.oAuth2.Get(url, model, &err)
	if e != nil {
		if err.Error.Message != "" {
			e.SetMessage(err.Error.Message)
		}

		return request, response, e
	}

	return request, response, nil
}

func (gd *GoogleDataStudio) patch(url string, requestBody []byte, model interface{}) (*http.Request, *http.Response, *errortools.Error) {
	err := google.ErrorResponse{}
	request, response, e := gd.oAuth2.Patch(url, bytes.NewBuffer(requestBody), model, &err)
	if e != nil {
		if err.Error.Message != "" {
			e.SetMessage(err.Error.Message)
		}

		return request, response, e
	}

	return request, response, nil
}

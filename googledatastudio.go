package googledatastudio

import (
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleDataStudio"
	apiURL  string = "https://datastudio.googleapis.com/v1"
)

// GoogleDataStudio stores GoogleDataStudio configuration
//
type GoogleDataStudio struct {
	Client *google.GoogleClient
}

// methods
//
func NewGoogleDataStudio(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) *GoogleDataStudio {
	config := google.GoogleClientConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	return &GoogleDataStudio{googleClient}
}

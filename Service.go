package googledatastudio

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
)

const (
	APIName string = "GoogleDataStudio"
	APIURL  string = "https://datastudio.googleapis.com/v1"
)

// Service stores Service configuration
//
type Service struct {
	googleService *google.Service
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
	Scope        string
}

// methods
//
func NewService(serviceConfig ServiceConfig, bigQueryService *bigquery.Service) *Service {
	config := google.ServiceConfig{
		APIName:      APIName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
		Scope:        serviceConfig.Scope,
	}

	googleService := google.NewService(config, bigQueryService)

	return &Service{googleService}
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", APIURL, path)
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}

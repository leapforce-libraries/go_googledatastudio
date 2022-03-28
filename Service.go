package googledatastudio

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleDataStudio"
	apiUrl  string = "https://datastudio.googleapis.com/v1"
)

type Service struct {
	googleService *google.Service
}

func NewServiceWithOAuth2(cfg *google.ServiceWithOAuth2Config) (*Service, *errortools.Error) {
	googleService, e := google.NewServiceWithOAuth2(cfg)
	if e != nil {
		return nil, e
	}
	return &Service{googleService}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.googleService.ApiKey()
}

func (service *Service) ApiCallCount() int64 {
	return service.googleService.ApiCallCount()
}

func (service *Service) ApiReset() {
	service.googleService.ApiReset()
}

func (service *Service) ErrorResponse() *google.ErrorResponse {
	return service.googleService.ErrorResponse()
}

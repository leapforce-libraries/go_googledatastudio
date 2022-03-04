package googledatastudio

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type AssetType string

const (
	AssetTypeReport     AssetType = "REPORT"
	AssetTypeDataSource AssetType = "DATA_SOURCE"
)

type AssetsResponse struct {
	Assets        []Asset `json:"assets"`
	NextPageToken string  `json:"nextPageToken"`
}

type Asset struct {
	Name             string    `json:"name"`
	Title            string    `json:"title"`
	AssetType        AssetType `json:"assetType"`
	Trashed          bool      `json:"trashed"`
	UpdateTime       time.Time `json:"updateTime"`
	UpdateByMeTime   time.Time `json:"updateByMeTime"`
	CreateTime       time.Time `json:"createTime"`
	LastViewByMeTime time.Time `json:"lastViewByMeTime"`
	Owner            string    `json:"owner"`
}

type SearchAssetsParams struct {
	Title          *string
	AssetTypes     AssetType
	IncludeTrashed *bool
	Owner          *Member
	OrderBy        *string
	PageSize       *int
	PageToken      *string
}

func (service *Service) SearchAssets(params *SearchAssetsParams) (*[]Asset, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("SearchAssetsParams cannot be nil.")
	}

	assets := []Asset{}

	query := []string{}

	query = append(query, fmt.Sprintf("assetTypes=%s", params.AssetTypes))
	if params.Title != nil {
		query = append(query, fmt.Sprintf("title=%s", *params.Title))
	}
	if params.IncludeTrashed != nil {
		query = append(query, fmt.Sprintf("includeTrashed=%v", *params.IncludeTrashed))
	}
	if params.Owner != nil {
		query = append(query, fmt.Sprintf("owner=%s", *params.Owner))
	}
	if params.OrderBy != nil {
		query = append(query, fmt.Sprintf("orderBy=%s", *params.OrderBy))
	}
	if params.PageSize != nil {
		query = append(query, fmt.Sprintf("pageSize=%v", *params.PageSize))
	}

	pageToken := params.PageToken

	for {
		if pageToken != nil {
			query = append(query, fmt.Sprintf("pageToken=%v", *params.PageToken))
		}

		assetsResponse := AssetsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("assets:search?%s", strings.Join(query, "&"))),
			ResponseModel: &assetsResponse,
		}
		_, _, e := service.googleService.HttpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		assets = append(assets, assetsResponse.Assets...)

		if params.PageToken != nil {
			break
		}
		if assetsResponse.NextPageToken == "" {
			break
		}

		pageToken = &assetsResponse.NextPageToken
	}

	return &assets, nil
}

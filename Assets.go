package googledatastudio

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
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

func (gd *GoogleDataStudio) SearchAssets(params *SearchAssetsParams) (*[]Asset, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("SearchAssetsParams cannot be nil.")
	}

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
	if params.PageToken != nil {
		query = append(query, fmt.Sprintf("pageToken=%v", *params.PageToken))
	}

	url := fmt.Sprintf("%s/assets:search?%s", apiURL, strings.Join(query, "&"))
	fmt.Println(url)

	assetsResponse := AssetsResponse{}

	_, _, e := gd.get(url, &assetsResponse)
	if e != nil {
		return nil, e
	}

	return &assetsResponse.Assets, nil
}

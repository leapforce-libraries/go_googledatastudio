package googledatastudio

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Role string

const (
	RoleEditor Role = "EDITOR"
	RoleOwner  Role = "OWNER"
	RoleViewer Role = "VIEWER"
)

type PermissionsObject struct {
	Permissions Permissions `json:"permissions"`
	Etag        string      `json:"etag"`
}

type Permissions struct {
	Editor Members `json:"EDITOR"`
	Owner  Members `json:"OWNER"`
	Viewer Members `json:"VIEWER"`
}

type Members struct {
	Members []Member `json:"members"`
}

type Member string

type GetPermissionsParams struct {
	AssetID string
	Role    *Role
}

func (gd *GoogleDataStudio) GetPermissions(params *GetPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("GetPermissionsParams cannot be nil.")
	}

	query := ""

	if params.Role != nil {
		query = fmt.Sprintf("?role=%s", *params.Role)
	}

	url := fmt.Sprintf("%s/assets/%s/permissions%s", apiURL, params.AssetID, query)
	fmt.Println(url)

	permissionsObject := PermissionsObject{}

	_, _, e := gd.Get(url, &permissionsObject)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

type PatchPermissionsParams struct {
	AssetID           string
	PermissionsObject *PermissionsObject
}

func (gd *GoogleDataStudio) PatchPermissions(params *PatchPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("GetPermissionsParams cannot be nil.")
	}
	if params.PermissionsObject == nil {
		return nil, errortools.ErrorMessage("PermissionsObject cannot be nil.")
	}

	url := fmt.Sprintf("%s/assets/%s/permissions", apiURL, params.AssetID)
	fmt.Println(url)

	requestBody := struct {
		Name        string            `json:"name"`
		Permissions PermissionsObject `json:"permissions"`
	}{
		params.AssetID,
		*params.PermissionsObject,
	}

	b, err := json.Marshal(&requestBody)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	permissionsObject := PermissionsObject{}

	_, _, e := gd.Patch(url, b, &permissionsObject)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

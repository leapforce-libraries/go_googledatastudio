package googledatastudio

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
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

func (service *Service) GetPermissions(params *GetPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("GetPermissionsParams cannot be nil.")
	}

	query := ""

	if params.Role != nil {
		query = fmt.Sprintf("?role=%s", *params.Role)
	}

	permissionsObject := PermissionsObject{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(fmt.Sprintf("assets/%s/permissions%s", params.AssetID, query)),
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.HTTPRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

type PatchPermissionsParams struct {
	AssetID           string
	PermissionsObject *PermissionsObject
}

func (service *Service) PatchPermissions(params *PatchPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("PatchPermissionsParams cannot be nil.")
	}
	if params.PermissionsObject == nil {
		return nil, errortools.ErrorMessage("PermissionsObject cannot be nil.")
	}

	requestBody := struct {
		Name        string            `json:"name"`
		Permissions PermissionsObject `json:"permissions"`
	}{
		params.AssetID,
		*params.PermissionsObject,
	}

	permissionsObject := PermissionsObject{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPatch,
		URL:           service.url(fmt.Sprintf("assets/%s/permissions", params.AssetID)),
		BodyModel:     requestBody,
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.HTTPRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

type AddMembersParams struct {
	AssetID string
	Role    Role
	Members *[]Member
}

func (service *Service) AddMembers(params *AddMembersParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("AddMembersParams cannot be nil.")
	}
	if params.Members == nil {
		return nil, errortools.ErrorMessage("Members cannot be nil.")
	}

	// add members in batches
	batchSize := 10
	members := []Member{}

	permissionsObject := PermissionsObject{}

	for i, member := range *params.Members {
		i++
		members = append(members, member)

		if i%batchSize == 0 || i == len(*params.Members) {

			requestBody := struct {
				Name    string   `json:"name"`
				Role    string   `json:"role"`
				Members []Member `json:"members"`
			}{
				params.AssetID,
				string(params.Role),
				members,
			}

			requestConfig := go_http.RequestConfig{
				Method:        http.MethodPost,
				URL:           service.url(fmt.Sprintf("assets/%s/permissions:addMembers", params.AssetID)),
				BodyModel:     requestBody,
				ResponseModel: &permissionsObject,
			}
			_, _, e := service.googleService.HTTPRequest(&requestConfig)
			if e != nil {
				return nil, e
			}

			members = []Member{}
		}
	}

	return &permissionsObject, nil
}

type RevokeAllPermissionsParams struct {
	AssetID string
	Members *[]Member
}

func (service *Service) RevokeAllPermissions(params *RevokeAllPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("RevokeAllPermissionsParams cannot be nil.")
	}
	if params.Members == nil {
		return nil, errortools.ErrorMessage("Members cannot be nil.")
	}

	// remove members in batches
	batchSize := 10
	members := []Member{}

	permissionsObject := PermissionsObject{}

	for i, member := range *params.Members {
		i++
		members = append(members, member)

		if i%batchSize == 0 || i == len(*params.Members) {
			requestBody := struct {
				Name    string   `json:"name"`
				Members []Member `json:"members"`
			}{
				params.AssetID,
				members,
			}

			requestConfig := go_http.RequestConfig{
				Method:        http.MethodPost,
				URL:           service.url(fmt.Sprintf("assets/%s/permissions:revokeAllPermissions", params.AssetID)),
				BodyModel:     requestBody,
				ResponseModel: &permissionsObject,
			}
			_, _, e := service.googleService.HTTPRequest(&requestConfig)
			if e != nil {
				return nil, e
			}
		}
	}

	return &permissionsObject, nil
}

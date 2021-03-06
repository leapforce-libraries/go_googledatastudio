package googledatastudio

import (
	"fmt"

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
		URL:           service.url(fmt.Sprintf("assets/%s/permissions%s", params.AssetID, query)),
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.Get(&requestConfig)
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
		URL:           service.url(fmt.Sprintf("assets/%s/permissions", params.AssetID)),
		BodyModel:     requestBody,
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.Patch(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

type AddMembersParams struct {
	AssetID string
	Role    Role
	Members *Members
}

func (service *Service) AddMembers(params *AddMembersParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("AddMembersParams cannot be nil.")
	}
	if params.Members == nil {
		return nil, errortools.ErrorMessage("Members cannot be nil.")
	}

	requestBody := struct {
		Name    string  `json:"name"`
		Role    string  `json:"role"`
		Members Members `json:"members"`
	}{
		params.AssetID,
		string(params.Role),
		*params.Members,
	}

	permissionsObject := PermissionsObject{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("assets/%s/permissions:addMembers", params.AssetID)),
		BodyModel:     requestBody,
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.Patch(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

type RevokeAllPermissionsParams struct {
	AssetID string
	Members *Members
}

func (service *Service) RevokeAllPermissions(params *RevokeAllPermissionsParams) (*PermissionsObject, *errortools.Error) {
	if params == nil {
		return nil, errortools.ErrorMessage("RevokeAllPermissionsParams cannot be nil.")
	}
	if params.Members == nil {
		return nil, errortools.ErrorMessage("Members cannot be nil.")
	}

	requestBody := struct {
		Name    string  `json:"name"`
		Members Members `json:"members"`
	}{
		params.AssetID,
		*params.Members,
	}

	permissionsObject := PermissionsObject{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("assets/%s/permissions:revokeAllPermissions", params.AssetID)),
		BodyModel:     requestBody,
		ResponseModel: &permissionsObject,
	}
	_, _, e := service.googleService.Post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &permissionsObject, nil
}

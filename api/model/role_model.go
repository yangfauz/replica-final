package model

import "replica-finalproject/api/entity"

//request
type CreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

//response
type GetRoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FormatGetAllRoleResponse(roles []entity.Role) []GetRoleResponse {
	rolesFormatter := []GetRoleResponse{}

	for _, role := range roles {
		roleFormatter := GetRoleResponse(role)
		rolesFormatter = append(rolesFormatter, roleFormatter)
	}

	return rolesFormatter
}

func FormatGetRoleResponse(role entity.Role) GetRoleResponse {

	roleFormatter := GetRoleResponse{}
	roleFormatter.ID = role.ID
	roleFormatter.Name = role.Name

	return roleFormatter
}

func FormatCreateRoleResponse(role entity.Role) CreateRoleResponse {

	roleFormatter := CreateRoleResponse{}
	roleFormatter.ID = role.ID
	roleFormatter.Name = role.Name

	return roleFormatter
}

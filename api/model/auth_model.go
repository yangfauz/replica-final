package model

import "replica-finalproject/api/entity"

//response
type LoginUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	RoleID   uint   `json:"role"`
	Token    string `json:"token"`
}

func FormatLoginUserResponse(user entity.User, token string) LoginUserResponse {

	userFormatter := LoginUserResponse{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.RoleID = user.RoleID
	userFormatter.Token = token

	return userFormatter
}

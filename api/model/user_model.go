package model

import "replica-finalproject/api/entity"

//request
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//response
type GetUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	RoleID   uint   `json:"role"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func FormatGetAllUserResponse(users []entity.User) []GetUserResponse {
	usersFormatter := []GetUserResponse{}

	for _, user := range users {
		userFormatter := GetUserResponse{}
		userFormatter.ID = user.ID
		userFormatter.Name = user.Name
		userFormatter.Username = user.Username
		userFormatter.RoleID = user.RoleID

		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

func FormatGetUserResponse(user entity.User) GetUserResponse {

	userFormatter := GetUserResponse{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.RoleID = user.RoleID

	return userFormatter
}

func FormatCreateUserResponse(user entity.User) GetUserResponse {

	userFormatter := GetUserResponse{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.RoleID = user.RoleID

	return userFormatter
}

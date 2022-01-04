package model

import "replica-finalproject/api/entity"

//response
type GetStatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FormatGetAllStatusResponse(statuses []entity.Status) []GetStatusResponse {
	statusesFormatter := []GetStatusResponse{}

	for _, status := range statuses {
		statusFormatter := GetStatusResponse(status)
		statusesFormatter = append(statusesFormatter, statusFormatter)
	}

	return statusesFormatter
}

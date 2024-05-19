package dto

import (
	"scheduling-app-back-end/internal/models"
)

type CreatePositionRequest struct {
	PositionName string `json:"position_name" gorm:"type:text"`
}

type GetPositionRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type PositionResponse struct {
	ID           int64                 `gorm:"type:bigint;primaryKey" json:"id"`
	PositionName string                `json:"position_name" gorm:"type:text"`
	Users        []*CreateUserResponse `gorm:"foreignKey:PositionID;references:ID" json:"users,omitempty"`
	Shifts       []*models.Shifts      `gorm:"foreignKey:PositionID;references:ID" json:"shifts,omitempty"`
	UsersArray   []int64               `gorm:"-" json:"users_array,omitempty"`
}

type PositionForUserCreateAndEdit struct {
	ID           int64  `json:"id"`
	PositionName string `json:"position_name" gorm:"type:text"`
}

func NewPositionResponse(positions *models.Positions) *PositionResponse {
	var allUsers []*CreateUserResponse
	for _, u := range positions.Users {
		i := NewUserResponse(u)
		allUsers = append(allUsers, i)
	}

	response := &PositionResponse{
		ID:           positions.ID,
		PositionName: positions.PositionName,
		Users:        allUsers,
		Shifts:       positions.Shifts,
		UsersArray:   positions.UsersArray,
	}

	return response
}

func PositionResponseFroUserAddEdit(positions *models.Positions) *PositionForUserCreateAndEdit {
	return &PositionForUserCreateAndEdit{ID: positions.ID, PositionName: positions.PositionName}
}

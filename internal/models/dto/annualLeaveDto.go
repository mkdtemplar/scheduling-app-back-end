package dto

import "scheduling-app-back-end/internal/models"

type CreateAnnualLeaveRequest struct {
	Email        string `gorm:"type:text" json:"email" binding:"required,email"`
	PositionName string `json:"position_name" gorm:"type:text"`
	StartDate    string `json:"start_date" db:"start_time" gorm:"type:time"`
	EndDate      string `json:"end_date" db:"end_time" gorm:"type:time"`
}

type CreateAnnualLeaveResponse struct {
	Email        string `gorm:"type:text" json:"email" binding:"required,email"`
	PositionName string `json:"position_name" gorm:"type:text"`
	StartDate    string `json:"start_date" db:"start_time" gorm:"type:time"`
	EndDate      string `json:"end_date" db:"end_time" gorm:"type:time"`
}

func NewAnnualLeaveResponse(leave *models.AnnualLeave) *CreateAnnualLeaveResponse {
	return &CreateAnnualLeaveResponse{
		Email:        leave.Email,
		PositionName: leave.PositionName,
		StartDate:    leave.StartDate,
		EndDate:      leave.EndDate,
	}
}

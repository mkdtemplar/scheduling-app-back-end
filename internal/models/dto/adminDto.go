package dto

import "scheduling-app-back-end/internal/models"

type CreateAdminRequest struct {
	ID       int64  `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	UserName string `gorm:"type:varchar(255);not null" json:"user_name" binding:"required,email"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}

type CreateAdminResponse struct {
	ID       int64  `gorm:"type:bigint;primaryKey" json:"id,string"`
	UserName string `gorm:"type:varchar(255);not null" json:"user_name"`
}

type AdminAuthorizationRequest struct {
	UserName string `gorm:"type:varchar(255);not null" json:"user_name" binding:"required,email"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}

func NewAdminResponse(newAdmin *models.Admin) *CreateAdminResponse {
	return &CreateAdminResponse{
		ID:       newAdmin.ID,
		UserName: newAdmin.UserName,
	}
}

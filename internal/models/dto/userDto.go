package dto

import (
	"scheduling-app-back-end/internal/models"
	"time"
)

type CreateUserRequest struct {
	ID           int64     `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	NameSurname  string    `gorm:"type:text" json:"name_surname" binding:"required"`
	Email        string    `gorm:"type:email" json:"email" binding:"required,email"`
	Password     string    `gorm:"type:text" json:"password" binding:"required"`
	PositionName string    `gorm:"type:text" json:"position_name" binding:"required"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"-"`
	UserID       int64     `gorm:"type:bigint" json:"user_id,string" binding:"required"`
}

type CreateUserResponse struct {
	ID           int64            `gorm:"type:bigint;primaryKey" json:"id,string"`
	NameSurname  string           `gorm:"type:text" json:"name_surname" binding:"required"`
	Email        string           `gorm:"type:email" json:"email" binding:"required,email"`
	PositionName string           `gorm:"type:text" json:"position_name" binding:"required"`
	Shifts       []*models.Shifts `gorm:"foreignKey:UserID;references:ID" json:"shifts,omitempty"`
	CreatedAt    time.Time        `gorm:"type:timestamp" json:"-"`
	UserID       int64            `gorm:"type:bigint" json:"user_id,string"`
}

func NewUserResponse(user *models.Users) *CreateUserResponse {
	return &CreateUserResponse{
		ID:           user.ID,
		NameSurname:  user.NameSurname,
		Email:        user.Email,
		PositionName: user.PositionName,
		Shifts:       user.Shifts,
		UserID:       user.UserID,
	}
}

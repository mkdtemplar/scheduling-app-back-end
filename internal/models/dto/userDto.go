package dto

import (
	"scheduling-app-back-end/internal/models"
	"time"
)

type CreateUserRequest struct {
	ID              int64       `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	FirstName       string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName        string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email           string      `gorm:"type:text" json:"email" binding:"required,email"`
	Password        string      `gorm:"type:text" json:"password" binding:"required"`
	CurrentPosition string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role            models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	CreatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID      int64       `gorm:"type:bigint" json:"position_id,string" binding:"required"`
}

type CreateUserResponse struct {
	ID                int64       `gorm:"type:bigint;primaryKey" json:"id,string"`
	FirstName         string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName          string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email             string      `gorm:"type:text" json:"email" binding:"required,email"`
	CurrentPosition   string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role              models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	PasswordChangedAt time.Time   `json:"password_changed_at,omitempty"`
	CreatedAt         time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID        int64       `gorm:"type:bigint" json:"position_id,string"`
}

func NewUserResponse(user *models.Users) *CreateUserResponse {
	return &CreateUserResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		CurrentPosition: user.CurrentPosition,
		Role:            user.Role,
		PositionID:      user.PositionID,
	}
}

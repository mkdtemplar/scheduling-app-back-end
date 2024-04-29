package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Users struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName       string    `gorm:"type:text" json:"first_name" binding:"required"`
	LastName        string    `gorm:"type:text" json:"last_name" binding:"required"`
	Email           string    `gorm:"type:text" json:"email" binding:"required,email"`
	Password        string    `gorm:"type:text" json:"password" binding:"required,min=8,max=32"`
	CurrentPosition string    `gorm:"type:text" json:"current_position" binding:"required"`
	Role            Role      `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	Shifts          []Shifts  `gorm:"foreignKey:UserID;references:ID" json:"shifts"`
	CreatedAt       time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time `gorm:"type:timestamp" json:"-"`
	PositionID      uuid.UUID `gorm:"type:uuid" json:"position_id"`
}

package models

import (
	"time"
)

type Users struct {
	ID                int64     `gorm:"type:bigint;primaryKey" json:"id"`
	NameSurname       string    `gorm:"type:text" json:"name_surname" binding:"required"`
	Email             string    `gorm:"type:text" json:"email" binding:"required,email"`
	Password          string    `gorm:"type:text" json:"password" binding:"omitempty,min=8,max=32"` // ✅ optional
	PositionName      string    `gorm:"type:text;index" json:"position_name" binding:"required"`
	Shifts            []*Shifts `gorm:"foreignKey:UserID;references:ID" json:"shifts,omitempty"`
	CreatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	UserID            int64     `json:"user_id" binding:"required"`
	UserPositionArray []int64   `gorm:"-" json:"user_position_array,omitempty"`
}

package models

import (
	"time"
)

type Users struct {
	ID                int64     `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	NameSurname       string    `gorm:"type:text" json:"name_surname" binding:"required"`
	Email             string    `gorm:"type:text" json:"email" binding:"required,email"`
	Password          string    `gorm:"type:text" json:"password" binding:"required,min=8,max=32"`
	PositionName      string    `gorm:"type:text" json:"position_name" binding:"required"`
	Shifts            []*Shifts `gorm:"foreignKey:UserID;references:ID" json:"shifts,omitempty"`
	CreatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	PositionID        int64     `json:"position_id,string" binding:"required"`
	UserPositionArray []int64   `gorm:"-" json:"user_position_array,omitempty"`
}

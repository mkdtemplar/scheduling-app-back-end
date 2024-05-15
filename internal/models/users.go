package models

import (
	"time"
)

type Role string

//const (
//	Admin Role = "admin"
//	Users  Role = "user"
//)

type Users struct {
	ID                int64     `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	FirstName         string    `gorm:"type:text" json:"first_name" binding:"required"`
	LastName          string    `gorm:"type:text" json:"last_name" binding:"required"`
	Email             string    `gorm:"type:text" json:"email" binding:"required,email"`
	Password          string    `gorm:"type:text" json:"password" binding:"required,min=8,max=32"`
	CurrentPosition   string    `gorm:"type:text" json:"current_position" binding:"required"`
	Role              Role      `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	Shifts            []*Shifts `gorm:"foreignKey:UserID;references:ID" json:"shifts,omitempty"`
	CreatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt         time.Time `gorm:"type:timestamp" json:"-"`
	PositionID        int64     `json:"position_id,string" binding:"required"`
	UserPositionArray []int64   `gorm:"-" json:"user_position_array,omitempty"`
}

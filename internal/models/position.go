package models

import (
	"time"
)

type Positions struct {
	ID           int64     `gorm:"type:bigint;primaryKey" json:"id"`
	PositionName string    `json:"position_name" gorm:"type:text"`
	Users        []*Users  `gorm:"foreignKey:PositionID;references:ID" json:"users,omitempty"`
	Shifts       []*Shifts `gorm:"foreignKey:PositionID;references:ID" json:"shifts,omitempty"`
	CreatedAt    time.Time `json:"-" gorm:"type:timestamp"`
	UpdatedAt    time.Time `json:"-" gorm:"type:timestamp"`
	UsersArray   []int64   `gorm:"-" json:"users_array,omitempty"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PositionName string     `json:"position_name" gorm:"type:text"`
	Employees    []Employee `json:"employees"`
	Shifts       []string   `json:"shifts" gorm:"type:text[];serializer:json"`
	StartTime    string     `json:"start_time" gorm:"type:varchar(10)"`
	EndTime      string     `json:"end_time" gorm:"type:varchar(10)"`
	CreatedAt    time.Time  `json:"-" gorm:"type:timestamp"`
	UpdatedAt    time.Time  `json:"-" gorm:"type:timestamp"`
}

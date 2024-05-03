package models

import (
	"time"

	"github.com/google/uuid"
)

type Positions struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PositionName string    `json:"position_name" gorm:"type:text"`
	Users        []Users   `gorm:"foreignKey:PositionID;references:ID" json:"users"`
	Shifts       []Shifts  `gorm:"foreignKey:PositionID;references:ID" json:"shifts"`
	CreatedAt    time.Time `json:"-" gorm:"type:timestamp"`
	UpdatedAt    time.Time `json:"-" gorm:"type:timestamp"`
}

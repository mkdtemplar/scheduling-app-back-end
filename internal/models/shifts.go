package models

import "github.com/google/uuid"

type Shifts struct {
	ID         uuid.UUID `json:"id" db:"id" gorm:"primaryKey;type:uuid"`
	Name       string    `json:"name" db:"name" gorm:"type:varchar(5)"`
	StartTime  string    `json:"start_time" db:"start_time" gorm:"type:time"`
	EndTime    string    `json:"end_time" db:"end_time" gorm:"type:time"`
	PositionID uuid.UUID `json:"position_id" db:"position_id" gorm:"type:uuid"`
	UserID     uuid.UUID `json:"user_id" db:"user_id" gorm:"type:uuid"`
}

package models

import (
	"github.com/jackc/pgx/pgtype"
)

type DailySchedule struct {
	ID             int64            `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	StartDate      string           `json:"start_date" db:"start_date" gorm:"type:time"`
	PositionsNames []string         `json:"positions_names" db:"positions" binding:"required" gorm:"type:text[]"`
	Employees      pgtype.TextArray `json:"employees" db:"employees" binding:"required" gorm:"type:text[]"`
	Shifts         pgtype.TextArray `json:"shifts" db:"shifts" binding:"required" gorm:"type:text[]"`
	Positions      []*Positions     `gorm:"foreignKey:PositionID;references:ID" json:"positions,omitempty"`
}

package models

import "github.com/lib/pq"

type DailySchedule struct {
	ID        int64  `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	StartDate string `json:"start_date" db:"start_date" gorm:"type:time"`

	PositionsNames pq.StringArray `json:"positions_names" db:"positions_names" gorm:"type:text[]"`
	Employees      pq.StringArray `json:"employees" db:"employees" gorm:"type:text[]"`
	Shifts         pq.StringArray `json:"shifts" db:"shifts" gorm:"type:text[]"`

	Positions []*Positions `gorm:"many2many:daily_schedule_positions;" json:"positions,omitempty"`
}

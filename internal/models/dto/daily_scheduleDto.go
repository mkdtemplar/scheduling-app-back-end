package dto

import "github.com/lib/pq"

type CreateDailyScheduleRequest struct {
	ID             int64          `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	StartDate      string         `json:"start_date" db:"start_date" gorm:"type:time"`
	PositionsNames pq.StringArray `json:"positions_names" db:"positions_names" binding:"required" gorm:"type:text[]"`
	Employees      pq.StringArray `json:"employees" db:"employees" binding:"required" gorm:"type:text[]"`
	Shifts         pq.StringArray `json:"shifts" db:"shifts" binding:"required" gorm:"type:text[]"`
}

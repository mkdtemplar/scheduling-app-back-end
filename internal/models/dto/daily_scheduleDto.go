package dto

type CreateDailyScheduleRequest struct {
	ID             int64    `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	StartDate      string   `json:"start_date" db:"start_date" gorm:"type:time"`
	PositionsNames []string `json:"positions_names" db:"positions" binding:"required" gorm:"type:text[]"`
	Employees      []string `json:"employees" db:"employees" binding:"required" gorm:"type:text[]"`
	Shifts         []string `json:"shifts" db:"shifts" binding:"required" gorm:"type:text[]"`
}

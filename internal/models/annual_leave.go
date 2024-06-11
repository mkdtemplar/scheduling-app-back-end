package models

type AnnualLeave struct {
	ID           int64  `json:"id,string" db:"id" gorm:"primaryKey;type:bigint"`
	Email        string `gorm:"type:text" json:"email" binding:"required,email"`
	PositionName string `json:"position_name" gorm:"type:text"`
	StartDate    string `json:"start_date" db:"start_time" gorm:"type:date"`
	EndDate      string `json:"end_date" db:"end_time" gorm:"type:date"`
}

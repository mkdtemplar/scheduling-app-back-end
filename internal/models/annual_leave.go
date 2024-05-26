package models

type AnnualLeave struct {
	ID        int64  `json:"id,string" db:"id" gorm:"primaryKey;type:bigint"`
	StartTime string `json:"start_time" db:"start_time" gorm:"type:time"`
	EndTime   string `json:"end_time" db:"end_time" gorm:"type:time"`
}

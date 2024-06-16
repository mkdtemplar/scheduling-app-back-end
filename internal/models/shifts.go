package models

type Shifts struct {
	ID         int64  `json:"id,string" db:"id" gorm:"primaryKey;type:bigint"`
	Name       string `json:"name" db:"name" gorm:"type:varchar(15)"`
	StartTime  string `json:"start_time" db:"start_time" gorm:"type:time"`
	EndTime    string `json:"end_time" db:"end_time" gorm:"type:time"`
	PositionID int64  `json:"position_id,string" db:"position_id" gorm:"type:bigint"`
	UserID     int64  `json:"user_id,string" db:"user_id" gorm:"type:bigint"`
}

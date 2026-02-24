package models

type Shifts struct {
	ID         int64  `json:"id,string" gorm:"primaryKey;type:bigint"`
	Name       string `json:"name" gorm:"type:varchar(15)"`
	StartTime  string `json:"start_time" gorm:"type:time"`
	EndTime    string `json:"end_time" gorm:"type:time"`
	PositionID int64  `json:"position_id,string" gorm:"type:bigint;index"`
	UserID     int64  `json:"user_id,string" gorm:"type:bigint;index"`
}

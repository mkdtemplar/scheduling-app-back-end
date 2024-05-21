package dto

type CreateShiftsRequest struct {
	ID         int64  `json:"id" db:"id" gorm:"primaryKey;type:bigint"`
	Name       string `json:"name" db:"name" gorm:"type:varchar(5)"`
	StartTime  string `json:"start_time" db:"start_time" gorm:"type:time"`
	EndTime    string `json:"end_time" db:"end_time" gorm:"type:time"`
	PositionID int64  `json:"position_id" db:"position_id" gorm:"type:bigint"`
	UserID     int64  `json:"user_id" db:"user_id" gorm:"type:bigint"`
}

package models

type Schedule struct {
	ID        int64        `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	StartDate string       `json:"start_date" db:"start_date" gorm:"type:time"`
	EndDate   string       `json:"end_date" db:"end_date" gorm:"type:time"`
	Positions []*Positions `gorm:"foreignKey:PositionID;references:ID" json:"positions,omitempty"`
}

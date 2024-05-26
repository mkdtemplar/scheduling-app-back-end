package models

type Positions struct {
	ID           int64     `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	PositionName string    `json:"position_name" gorm:"type:text"`
	Users        []*Users  `gorm:"foreignKey:UserID;references:ID" json:"users,omitempty"`
	Shifts       []*Shifts `gorm:"foreignKey:PositionID;references:ID" json:"shifts,omitempty"`
	UsersArray   []int64   `gorm:"-" json:"users_array,omitempty"`
	//PositionID   int64     `json:"position_id,string" db:"position_id" gorm:"type:bigint"`
}

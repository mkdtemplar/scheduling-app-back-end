package models

type Positions struct {
	ID           int64  `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	PositionName string `json:"position_name" gorm:"type:text"`

	// Associations
	Users  []*Users  `gorm:"foreignKey:UserID;references:ID" json:"users,omitempty"`
	Shifts []*Shifts `gorm:"foreignKey:PositionID;references:ID" json:"shifts,omitempty"`

	// This is not a foreign key, just a helper for incoming JSON (ignored by GORM)
	UsersArray []int64 `gorm:"-" json:"users_array,omitempty"`

	// Remove this line to avoid foreign key confusion:
	// PositionID   int64   `json:"position_id,string" binding:"required"`
}

package models

type Admin struct {
	ID       int64  `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	UserName string `gorm:"type:varchar(255);not null" json:"user_name" binding:"required"`
}

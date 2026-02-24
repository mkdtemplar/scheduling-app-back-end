package models

import "time"

type DailyAssignment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SenderEmail  string    `gorm:"type:text;index" json:"sender_email"`
	PositionName string    `gorm:"type:text;index" json:"position_name"`
	Description  string    `gorm:"type:text" json:"description"`
	SentCount    int       `gorm:"type:int" json:"sent_count"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"created_at"`
}

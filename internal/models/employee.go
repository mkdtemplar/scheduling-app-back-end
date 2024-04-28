package models

import "github.com/google/uuid"

type Employee struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName       string    `gorm:"type:text" json:"first_name"`
	LastName        string    `gorm:"type:text" json:"last_name"`
	HashedPassword  string    `gorm:"type:text" json:"hashed_password"`
	CurrentPosition string    `gorm:"type:text" json:"current_position"`
}

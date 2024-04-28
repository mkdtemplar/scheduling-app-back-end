package models

import "github.com/google/uuid"

type Administrator struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Username string    `json:"username" gorm:"varchar(25);unique" binding:"required,email"`
	Password string    `json:"password" gorm:"varchar(32)" binding:"required,min=8,max=32"`
}

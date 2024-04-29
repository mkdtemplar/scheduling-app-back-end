package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Authorization struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type JwtUser struct {
	ID        uuid.UUID `json:"id" db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"first_name" db:"first_name" gorm:"type:varchar(55)"`
	LastName  string    `json:"last_name" db:"last_name" gorm:"type:varchar(55)"`
}

type TokenPairs struct {
	Token        string `json:"access_token" db:"access_token" gorm:"type:varchar(255)"`
	RefreshToken string `json:"refresh_token" db:"refresh_token" gorm:"type:varchar(255)"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Authorization) GenerateTokenPairs(user *JwtUser) (TokenPairs, error) {
	return TokenPairs{}, nil
}

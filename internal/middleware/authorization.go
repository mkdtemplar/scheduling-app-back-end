package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Authorization struct {
	Issuer        string
	Audience      string
	JWTSecret     string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

func NewAuthorization(issuer string, audience string, secret string, tokenExpiry time.Duration,
	refreshExpiry time.Duration, cookieDomain string, cookiePath string, cookieName string) IJWTInterfaces {
	return &Authorization{
		Issuer:        issuer,
		Audience:      audience,
		JWTSecret:     secret,
		TokenExpiry:   tokenExpiry,
		RefreshExpiry: refreshExpiry,
		CookieDomain:  cookieDomain,
		CookiePath:    cookiePath,
		CookieName:    cookieName,
	}
}

type JwtUser struct {
	ID        uuid.UUID `json:"id" db:"id" gorm:"type:uuid"`
	FirstName string    `json:"first_name" db:"first_name" gorm:"type:varchar(55)"`
	LastName  string    `json:"last_name" db:"last_name" gorm:"type:varchar(55)"`
	Email     string    `json:"email" db:"email" gorm:"type:varchar(255)"`
}

type TokenPairs struct {
	Token        string `json:"access_token" db:"access_token" gorm:"type:varchar(255)"`
	RefreshToken string `json:"refresh_token" db:"refresh_token" gorm:"type:varchar(255)"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Authorization) GenerateTokenPairs(user *JwtUser) (TokenPairs, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = user.ID.String()
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	signedAccessToken, err := token.SignedString([]byte(j.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = user.ID.String()
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	refreshTokenClaims["ext"] = time.Now().Add(j.RefreshExpiry).UTC().Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(j.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	tokenPairs := TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}
	return tokenPairs, nil
}

func (j *Authorization) GetRefreshCookie(refreshToken string, ctx *gin.Context) {
	ctx.SetCookie(j.CookieName, refreshToken, int(j.RefreshExpiry.Seconds()), j.CookiePath, j.CookieDomain, true, true)
}

func (j *Authorization) GetExpiredRefreshCookie(ctx *gin.Context) {
	ctx.SetCookie(j.CookieName, "", 0, j.CookiePath, j.CookieDomain, true, true)
}

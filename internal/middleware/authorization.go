package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"scheduling-app-back-end/internal/repository/db"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	ID        int64  `json:"id" db:"id" gorm:"type:bigserial"`
	FirstName string `gorm:"type:text" json:"first_name" binding:"required"`
	LastName  string `gorm:"type:text" json:"last_name" binding:"required"`
	Email     string `json:"email" db:"email" gorm:"type:varchar(255)"`
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
	claims["name"] = fmt.Sprintf("%s%s", user.FirstName, user.LastName)
	claims["sub"] = strconv.Itoa(int(user.ID))
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
	refreshTokenClaims["sub"] = strconv.Itoa(int(user.ID))
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
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie(j.CookieName, refreshToken, int(j.RefreshExpiry.Seconds()), j.CookiePath, j.CookieDomain, true, true)
}

func (j *Authorization) GetExpiredRefreshCookie(ctx *gin.Context) {

	ctx.SetCookie(j.CookieName, "", 0, j.CookiePath, j.CookieDomain, true, true)
}

func (j *Authorization) RefreshToken(ctx *gin.Context) {
	userRepo := db.NewUserRepo()
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == j.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(j.JWTSecret), nil
			})
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("unauthorized")})
				return
			}

			userId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("cannot parse id")})
			}
			user, err := userRepo.GetUserById(ctx, int64(userId))
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("user not found")})
				return
			}
			log.Println(user)

			u := &JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			}

			tokenPairs, err := j.GenerateTokenPairs(u)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("cannot generate tokens")})
				return
			}

			ctx.SetSameSite(http.SameSiteStrictMode)
			ctx.SetCookie(j.CookieName, tokenPairs.RefreshToken, int(j.RefreshExpiry.Seconds()), j.CookiePath, j.CookieDomain, true, true)

			ctx.JSON(http.StatusOK, tokenPairs)

		}
	}
}

func (j *Authorization) Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie(j.CookieName, "", -1, j.CookiePath, j.CookieDomain, true, true)
	ctx.Writer.WriteHeader(http.StatusAccepted)
}

func (j *Authorization) GetTokenFromHeaderAndVerify(ctx *gin.Context) (string, *Claims, error) {
	ctx.Writer.Header().Add("Vary", "Authorization")

	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil, errors.New("authorization header is missing")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("authorization header is invalid")
	}

	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("authorization header is invalid")
	}

	token := headerParts[1]

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.JWTSecret), nil
	})
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("token is expired")
		}
		return "", nil, err
	}

	if claims.Issuer != j.Issuer {
		return "", nil, errors.New("invalid issuer")
	}

	return token, claims, nil
}

func (j *Authorization) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, err := j.GetTokenFromHeaderAndVerify(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

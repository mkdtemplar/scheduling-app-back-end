package middleware

import (
	"errors"
	"fmt"
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
	ID       int64  `json:"id" binding:"required"`
	Username string `json:"user_name" binding:"required"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"-"`
}

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (j *Authorization) SetRefreshCookie(ctx *gin.Context, refreshToken string) {
	secure := true
	sameSite := http.SameSiteStrictMode

	if j.CookieDomain == "" || strings.Contains(j.CookieDomain, "localhost") {
		secure = false
		sameSite = http.SameSiteLaxMode
	}

	ctx.SetSameSite(sameSite)
	ctx.SetCookie(
		j.CookieName,
		refreshToken,
		int(j.RefreshExpiry.Seconds()),
		j.CookiePath,
		j.CookieDomain,
		secure,
		true,
	)
}

func (j *Authorization) GenerateTokenPairs(user *JwtUser) (TokenPairs, error) {
	access := jwt.New(jwt.SigningMethodHS256)
	accessClaims := access.Claims.(jwt.MapClaims)

	accessClaims["name"] = user.Username
	accessClaims["sub"] = strconv.Itoa(int(user.ID))
	accessClaims["aud"] = j.Audience
	accessClaims["iss"] = j.Issuer
	accessClaims["iat"] = time.Now().UTC().Unix()
	accessClaims["typ"] = "JWT"
	accessClaims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	signedAccessToken, err := access.SignedString([]byte(j.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refresh.Claims.(jwt.MapClaims)

	refreshClaims["sub"] = strconv.Itoa(int(user.ID))
	refreshClaims["iat"] = time.Now().UTC().Unix()

	refreshClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	signedRefreshToken, err := refresh.SignedString([]byte(j.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (j *Authorization) GetRefreshCookie(refreshToken string, ctx *gin.Context) {
	secure := true
	sameSite := http.SameSiteStrictMode

	if j.CookieDomain == "" || strings.Contains(j.CookieDomain, "localhost") {
		secure = false
		sameSite = http.SameSiteLaxMode
	}

	ctx.SetSameSite(sameSite)
	ctx.SetCookie(
		j.CookieName,
		refreshToken,
		int(j.RefreshExpiry.Seconds()),
		j.CookiePath,
		j.CookieDomain,
		secure,
		true,
	)
}

func (j *Authorization) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie(j.CookieName)
	if err != nil || strings.TrimSpace(refreshToken) == "" {
		// ✅ Always respond (no silent return)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh cookie"})
		return
	}

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JWTSecret), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("unauthorized").Error()})
		return
	}

	adminID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse id"})
		return
	}

	adminRepo := db.NewAdminRepo()
	admin, err := adminRepo.GetAdminById(ctx, int64(adminID))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	u := &JwtUser{ID: admin.ID, Username: admin.UserName}
	tokenPairs, err := j.GenerateTokenPairs(u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate tokens"})
		return
	}

	j.GetRefreshCookie(tokenPairs.RefreshToken, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokenPairs.Token,
	})
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
		_, claims, err := j.GetTokenFromHeaderAndVerify(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("admin_mail", strings.TrimSpace(claims.Name))
		c.Set("admin_id", strings.TrimSpace(claims.Subject))
		c.Next()
	}
}

func (j *Authorization) AuthStatus(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie(j.CookieName)
	if err != nil || strings.TrimSpace(refreshToken) == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"isLoggedIn": false})
		return
	}

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JWTSecret), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"isLoggedIn": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
		"userId":     claims.Subject,
	})
}

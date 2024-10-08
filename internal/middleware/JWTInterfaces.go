package middleware

import (
	"github.com/gin-gonic/gin"
)

type IJWTInterfaces interface {
	GenerateTokenPairs(user *JwtUser) (TokenPairs, error)
	GetRefreshCookie(refreshToken string, ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetTokenFromHeaderAndVerify(ctx *gin.Context) (string, *Claims, error)
	AuthRequired() gin.HandlerFunc
}

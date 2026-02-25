package api

import (
	"errors"
	"net/http"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (adm *AdminHandler) Authorization(ctx *gin.Context) {
	var requestPayload struct {
		Username string `json:"user_name" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminByEmail, err := adm.IAdminInterfaces.GetAdminByEmail(ctx, requestPayload.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	valid, err := utils.CheckPassword(requestPayload.Password, adminByEmail.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	user := middleware.JwtUser{
		ID:       adminByEmail.ID,
		Username: adminByEmail.UserName,
	}

	tokens, err := adm.IJWTInterfaces.GenerateTokenPairs(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate tokens"})
		return
	}

	adm.IJWTInterfaces.GetRefreshCookie(tokens.RefreshToken, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokens.Token,
	})
}

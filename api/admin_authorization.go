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
		Username string `json:"user_name" binding:"required" gorm:"type:email"`
		Password string `json:"password" binding:"required" gorm:"type:password"`
	}

	if err := ctx.ShouldBindJSON(&requestPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if requestPayload.Username == "" || requestPayload.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Username & Password required"})
		return
	}

	adminByEmail, err := adm.IAdminInterfaces.GetAdminByEmail(ctx, requestPayload.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	valid, err := utils.CheckPassword(requestPayload.Password, adminByEmail.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	testUser := middleware.JwtUser{
		ID:       adminByEmail.ID,
		Username: adminByEmail.UserName,
	}

	tokens, err := adm.IJWTInterfaces.GenerateTokenPairs(&testUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
		return
	}

	adm.IJWTInterfaces.GetRefreshCookie(tokens.Token, ctx)

	ctx.JSON(http.StatusAccepted, tokens)

}

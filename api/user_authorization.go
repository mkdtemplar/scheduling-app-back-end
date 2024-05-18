package api

import (
	"errors"
	"net/http"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (usr *UserHandler) Authorization(ctx *gin.Context) {

	var requestPayload struct {
		Email    string `json:"email" binding:"required" gorm:"type:email"`
		Password string `json:"password" binding:"required" gorm:"type:password"`
	}

	if err := ctx.ShouldBindJSON(&requestPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userFromDb, err := usr.IUserRepository.GetUserByEmail(ctx, requestPayload.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	valid, err := utils.CheckPassword(requestPayload.Password, userFromDb.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	testUser := middleware.JwtUser{
		ID:          userFromDb.ID,
		NameSurname: userFromDb.NameSurname,
	}

	tokens, err := usr.IJWTInterfaces.GenerateTokenPairs(&testUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
		return
	}

	usr.IJWTInterfaces.GetRefreshCookie(tokens.Token, ctx)

	ctx.JSON(http.StatusAccepted, tokens)

}

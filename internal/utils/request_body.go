package utils

import (
	"encoding/json"
	"io"
	"scheduling-app-back-end/internal/models"

	"github.com/gin-gonic/gin"
)

func ParseUserPrefRequestBody(ctx *gin.Context) (*models.Users, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	user := &models.Users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return &models.Users{}, err
	}

	return user, nil
}

func ParseAdminRequestBody(ctx *gin.Context) (*models.Admin, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	admin := &models.Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		return &models.Admin{}, err
	}
	return admin, nil
}

func ParseShiftRequestBody(ctx *gin.Context) (*models.Shifts, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	shift := &models.Shifts{}
	err = json.Unmarshal(body, &shift)
	if err != nil {
		return &models.Shifts{}, err
	}
	return shift, nil
}

func ParsePositionRequestBody(ctx *gin.Context) (*models.Positions, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	positions := &models.Positions{}
	err = json.Unmarshal(body, &positions)
	if err != nil {
		return &models.Positions{}, err
	}
	return positions, nil
}

func ParseAdminAuthorizationBody(ctx *gin.Context) (*models.Admin, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	admin := &models.Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		return &models.Admin{}, err
	}
	return admin, nil
}

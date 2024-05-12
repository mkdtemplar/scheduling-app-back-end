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

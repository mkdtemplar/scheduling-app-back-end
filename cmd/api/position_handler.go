package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPositionRequest struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	PositionName string            `json:"position_name" gorm:"type:text"`
	Employees    []models.Employee `json:"employees" gorm:"type:text[];serializer:json"`
	Shifts       []string          `json:"shifts" gorm:"type:text[];serializer:json"`
	StartTime    string            `json:"start_time" gorm:"type:varchar(10)"`
	EndTime      string            `json:"end_time" gorm:"type:varchar(10)"`
	CreatedAt    time.Time         `json:"-" gorm:"type:timestamp"`
	UpdatedAt    time.Time         `json:"-" gorm:"type:timestamp"`
}

func NewPositionHandler(IPositionRepository interfaces.IPositionsRepository) *PositionHandler {
	return &PositionHandler{
		IPositionsRepository: IPositionRepository,
	}
}

func (i *PositionHandler) CreatePosition(ctx *gin.Context) {
	var req *createPositionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &models.Position{
		ID:           utils.GenerateID(),
		PositionName: req.PositionName,
		Shifts:       req.Shifts,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	newPosition, err := i.IPositionsRepository.CreatePosition(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newPosition)
}

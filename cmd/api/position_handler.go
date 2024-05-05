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
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PositionName string    `json:"position_name" gorm:"type:text"`
	StartTime    string    `json:"start_time" gorm:"type:varchar(10)"`
	UpdatedAt    time.Time `json:"-" gorm:"type:timestamp"`
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

	arg := &models.Positions{
		ID:           utils.GenerateID(),
		PositionName: req.PositionName,
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

func (i *PositionHandler) AllPositions(ctx *gin.Context) {
	positions, err := i.IPositionsRepository.AllPositions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, positions)
}

func (i *PositionHandler) GetPositionById(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Params.ByName("id"))

	position, err := i.IPositionsRepository.GetPositionByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, position)
}

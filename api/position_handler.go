package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewPositionHandler(IPositionRepository interfaces.IPositionsRepository) *PositionHandler {
	return &PositionHandler{
		IPositionsRepository: IPositionRepository,
	}
}

func (i *PositionHandler) CreatePosition(ctx *gin.Context) {
	var req *dto.CreatePositionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &models.Positions{
		ID:           req.ID,
		PositionName: req.PositionName,
	}

	newPosition, err := i.IPositionsRepository.CreatePosition(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newPosition)
}

func (i *PositionHandler) AllPositions(ctx *gin.Context) {
	var allPositions []*dto.PositionResponse
	positions, err := i.IPositionsRepository.AllPositions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, position := range positions {
		position.UsersArray = append(position.UsersArray, position.ID)
	}
	for _, position := range positions {
		allPositions = append(allPositions, dto.NewPositionResponse(position))
	}

	ctx.JSON(http.StatusOK, allPositions)
}

func (i *PositionHandler) GetPositionById(ctx *gin.Context) {
	var req dto.GetPositionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	position, err := i.IPositionsRepository.GetPositionByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

func (i *PositionHandler) GetPositionByIdForEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	position, err := i.IPositionsRepository.GetPositionByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

func (i *PositionHandler) AllPositionsForUserAddEdit(ctx *gin.Context) {
	var allPositions []*dto.PositionResponse

	positions, err := i.IPositionsRepository.AllPositionsForUserAddEdit(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, position := range positions {
		allPositions = append(allPositions, dto.NewPositionResponse(position))
	}

	ctx.JSON(http.StatusOK, allPositions)
}

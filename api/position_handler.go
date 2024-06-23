package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
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

func (i *PositionHandler) AllPositionsForDailySchedule(ctx *gin.Context) {
	var allPositionsForDailySchedules []*dto.PositionForDailyScheduleResponse
	allPositions, err := i.IPositionsRepository.AllPositionsForDailySchedule(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, position := range allPositions {
		allPositionsForDailySchedules = append(allPositionsForDailySchedules, dto.NewPositionForDailyScheduleResponse(position))
	}

	ctx.JSON(http.StatusOK, allPositionsForDailySchedules)
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

func (i *PositionHandler) UpdatePosition(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	positionFromDb, err := i.IPositionsRepository.GetPositionByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	positionForEdit, err := utils.ParsePositionRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	positionFromDb, err = i.IPositionsRepository.UpdatePosition(ctx, positionFromDb.ID, positionForEdit.ID, positionForEdit.PositionName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, positionFromDb)
}

func (i *PositionHandler) DeletePosition(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err = i.IPositionsRepository.DeletePosition(ctx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"error": false, "message": "position deleted"})
}

package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"

	"github.com/gin-gonic/gin"
)

func NewShiftsHandler(IShiftsRepository interfaces.IShiftsInterfaces) *ShiftsHandler {
	return &ShiftsHandler{
		IShiftsInterfaces: IShiftsRepository,
	}
}

func (sh *ShiftsHandler) CreateShift(ctx *gin.Context) {
	var req *models.Shifts

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &models.Shifts{
		ID:         req.ID,
		Name:       req.Name,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		PositionID: req.PositionID,
		UserID:     req.UserID,
	}

	newShift, err := sh.IShiftsInterfaces.CreateShifts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, newShift)
}

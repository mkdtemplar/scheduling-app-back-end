package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"strconv"

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

func (sh *ShiftsHandler) GetShiftById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shift, err := sh.IShiftsInterfaces.GetShiftById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shift)
}

func (sh *ShiftsHandler) GetShiftByName(ctx *gin.Context) {

	shift, err := sh.IShiftsInterfaces.GetShiftByName(ctx, ctx.Query("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shift)
}

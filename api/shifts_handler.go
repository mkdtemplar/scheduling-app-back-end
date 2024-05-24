package api

import (
	"errors"
	"fmt"
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	shift, err := sh.IShiftsInterfaces.GetShiftByName(ctx, ctx.Params.ByName("name"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shift)
}

func (sh *ShiftsHandler) GetAllShifts(ctx *gin.Context) {
	var shifts []*models.Shifts

	shifts, err := sh.IShiftsInterfaces.GetAllShifts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shifts)
}

func (sh *ShiftsHandler) UpdateShift(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shiftFromDb, err := sh.IShiftsInterfaces.GetShiftById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	shiftForEdit, err := utils.ParseShiftRequestBody(ctx)
	fmt.Println(shiftForEdit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shiftFromDb, err = sh.IShiftsInterfaces.UpdateShift(ctx, shiftFromDb.ID, shiftForEdit.ID, shiftForEdit.Name, shiftForEdit.StartTime,
		shiftForEdit.EndTime, shiftForEdit.PositionID, shiftForEdit.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, shiftFromDb)
}

func (sh *ShiftsHandler) DeleteShift(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = sh.DB.Delete(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"error": false, "message": "shift deleted"})
}

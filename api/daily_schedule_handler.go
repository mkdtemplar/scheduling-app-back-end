package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDailyScheduleHandler(IDailyScheduleInterfaces interfaces.IDailyScheduleInterfaces) *DailyScheduleHandlers {
	return &DailyScheduleHandlers{
		IDailyScheduleInterfaces: IDailyScheduleInterfaces,
	}
}

func (d *DailyScheduleHandlers) CreateDailySchedule(ctx *gin.Context) {
	var req *dto.CreateDailyScheduleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := &models.DailySchedule{
		ID:             req.ID,
		StartDate:      req.StartDate,
		PositionsNames: req.PositionsNames,
		Employees:      req.Employees,
		Shifts:         req.Shifts,
	}

	newDailySchedule, err := d.IDailyScheduleInterfaces.CrateDailySchedule(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, newDailySchedule)
}

func (d *DailyScheduleHandlers) GetAllDailySchedules(ctx *gin.Context) {

	allDailySchedules, err := d.IDailyScheduleInterfaces.GetAllDailySchedules(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, allDailySchedules)
}

func (d *DailyScheduleHandlers) GetDailyScheduleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dailySchedule, err := d.IDailyScheduleInterfaces.GetDailyScheduleById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dailySchedule)
}

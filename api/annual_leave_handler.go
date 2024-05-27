package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"

	"github.com/gin-gonic/gin"
)

func NewAnnualLeaveHandler(IAnnualLeaveInterfaces interfaces.IAnnualLeaveInterfaces) *AnnualLeaveHandler {
	return &AnnualLeaveHandler{
		IAnnualLeaveInterfaces: IAnnualLeaveInterfaces,
	}
}

func (a *AnnualLeaveHandler) CreateAnnualLeave(ctx *gin.Context) {
	var req *dto.CreateAnnualLeaveRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &models.AnnualLeave{
		Email:        req.Email,
		PositionName: req.PositionName,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}

	newAnnualLeave, err := a.IAnnualLeaveInterfaces.CreateAnnualLeave(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewAnnualLeaveResponse(newAnnualLeave)
	ctx.JSON(http.StatusOK, response)
}

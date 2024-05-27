package api

import (
	"fmt"
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"

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

	htmlMessage := fmt.Sprintf(`
		<strong>Annual Leave Request Send for Approval</strong><br>
		This email confirms that user with email %s send request for annual leave from %s until %s
`, response.Email, response.StartDate, response.EndDate)

	msg := models.MailData{
		To:      response.Email,
		From:    "admin@example.com",
		Subject: "Annual Leave Confirmation",
		Content: htmlMessage,
	}

	utils.SendMsg(msg)

	msg = models.MailData{
		To:      "admin@example.com",
		From:    response.Email,
		Subject: "Annual Leave Confirmation",
		Content: htmlMessage,
	}
	utils.SendMsg(msg)
}

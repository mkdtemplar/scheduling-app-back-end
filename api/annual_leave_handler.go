package api

import "scheduling-app-back-end/internal/repository/interfaces"

func NewAnnualLeaveHandler(IAnnualLeaveRepository interfaces.IAnnualLeaveRepository) *AnnualLeaveHandler {
	return &AnnualLeaveHandler{
		IAnnualLeaveRepository: IAnnualLeaveRepository,
	}
}

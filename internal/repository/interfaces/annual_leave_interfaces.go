package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IAnnualLeaveInterfaces interface {
	CreateAnnualLeave(ctx context.Context, annualLeave *models.AnnualLeave) (*models.AnnualLeave, error)
}

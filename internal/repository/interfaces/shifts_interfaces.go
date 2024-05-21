package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IShiftsInterfaces interface {
	CreateShifts(ctx context.Context, shift *models.Shifts) (*models.Shifts, error)
}

package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IShiftsInterfaces interface {
	CreateShifts(ctx context.Context, shift *models.Shifts) (*models.Shifts, error)
	GetAllShifts(ctx context.Context) ([]*models.Shifts, error)
	GetShiftById(ctx context.Context, id int64) (*models.Shifts, error)
	GetShiftByName(ctx context.Context, shiftName string) (*models.Shifts, error)
}

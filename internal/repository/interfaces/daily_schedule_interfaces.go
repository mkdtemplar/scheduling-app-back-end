package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IDailyScheduleInterfaces interface {
	CrateDailySchedule(ctx context.Context, dailySchedule *models.DailySchedule) (*models.DailySchedule, error)
}

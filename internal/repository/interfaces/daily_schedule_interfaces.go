package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IDailyScheduleInterfaces interface {
	CrateDailySchedule(ctx context.Context, dailySchedule *models.DailySchedule) (*models.DailySchedule, error)
	GetAllDailySchedules(ctx context.Context) ([]*models.DailySchedule, error)
	GetDailyScheduleById(ctx context.Context, id int64) (*models.DailySchedule, error)
}

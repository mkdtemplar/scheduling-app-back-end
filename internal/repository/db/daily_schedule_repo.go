package db

import (
	"context"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewDailyScheduleRepo() interfaces.IDailyScheduleInterfaces {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CrateDailySchedule(ctx context.Context, dailySchedule *models.DailySchedule) (*models.DailySchedule, error) {
	if dailySchedule == nil {
		return &models.DailySchedule{}, nil
	}

	if err := p.DB.WithContext(ctx).Model(&models.DailySchedule{}).Create(&dailySchedule).Error; err != nil {
		return &models.DailySchedule{}, err
	}
	return dailySchedule, nil
}

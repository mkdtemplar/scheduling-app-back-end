package db

import (
	"context"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"strings"
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

func (p *PostgresDB) GetAllDailySchedules(ctx context.Context) ([]*models.DailySchedule, error) {
	var dailySchedules []*models.DailySchedule
	err := p.DB.WithContext(ctx).Model(&models.DailySchedule{}).Find(&dailySchedules).Error
	if err != nil {
		return nil, err
	}
	for _, dailySchedule := range dailySchedules {
		dailySchedule.StartDate = dailySchedule.StartDate[:strings.IndexByte(dailySchedule.StartDate, 'T')]
	}
	return dailySchedules, nil
}

func (p *PostgresDB) GetDailyScheduleById(ctx context.Context, id int64) (*models.DailySchedule, error) {
	var dailySchedule *models.DailySchedule

	if err := p.DB.WithContext(ctx).Model(&models.DailySchedule{}).Where("id = ?", id).Take(&dailySchedule).Error; err != nil {
		return &models.DailySchedule{}, err
	}

	dailyScheduleFind := &models.DailySchedule{
		ID:             dailySchedule.ID,
		StartDate:      dailySchedule.StartDate[:strings.IndexByte(dailySchedule.StartDate, 'T')],
		PositionsNames: dailySchedule.PositionsNames,
		Employees:      dailySchedule.Employees,
		Shifts:         dailySchedule.Shifts,
		Positions:      dailySchedule.Positions,
	}
	return dailyScheduleFind, nil
}

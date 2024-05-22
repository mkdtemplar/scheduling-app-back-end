package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewShiftsRepo() interfaces.IShiftsInterfaces {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateShifts(ctx context.Context, shift *models.Shifts) (*models.Shifts, error) {
	if shift == nil {
		return nil, errors.New("shifts is empty")
	}

	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Create(&shift).Error; err != nil {
		return nil, err
	}

	return shift, nil
}

func (p *PostgresDB) GetAllShifts(ctx context.Context) ([]*models.Shifts, error) {
	var shifts []*models.Shifts

	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Find(&shifts).Error; err != nil {
		return nil, err
	}
	return shifts, nil
}

func (p *PostgresDB) GetShiftById(ctx context.Context, id int64) (*models.Shifts, error) {
	shift := &models.Shifts{}
	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Take(&shift, id).Error; err != nil {
		return nil, err
	}
	return shift, nil
}

func (p *PostgresDB) GetShiftByName(ctx context.Context, name string) (*models.Shifts, error) {
	shift := &models.Shifts{}
	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Take(&shift, name).Error; err != nil {
		return nil, err
	}
	return shift, nil
}

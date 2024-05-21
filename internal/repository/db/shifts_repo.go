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

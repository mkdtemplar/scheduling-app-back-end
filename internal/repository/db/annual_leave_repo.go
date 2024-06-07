package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewAnnualLeaveRepo() interfaces.IAnnualLeaveInterfaces {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateAnnualLeave(ctx context.Context, annualLeave *models.AnnualLeave) (*models.AnnualLeave, error) {
	if annualLeave == nil {
		return nil, errors.New("AnnualLeave is empty")
	}

	if err := p.DB.WithContext(ctx).Model(&models.AnnualLeave{}).Create(&annualLeave).Error; err != nil {
		return nil, err
	}
	return annualLeave, nil
}

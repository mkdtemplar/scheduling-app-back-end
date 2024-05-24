package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"

	"gorm.io/gorm"
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
	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Where("id = ?", id).First(&shift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return shift, nil
}

func (p *PostgresDB) GetShiftByName(ctx context.Context, name string) (*models.Shifts, error) {
	shift := &models.Shifts{}
	if err := p.DB.WithContext(ctx).Model(&models.Shifts{}).Where("name = ?", name).First(&shift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return shift, nil
}

func (p *PostgresDB) UpdateShift(ctx context.Context, id int64, idUpdated int64, name string, startTime string,
	endTime string, positionID int64, userId int64) (*models.Shifts, error) {

	shiftForUpdate, err := p.GetShiftById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = p.DB.WithContext(ctx).Model(shiftForUpdate).Where("id = ?", id).
		UpdateColumns(map[string]interface{}{"id": idUpdated, "name": name, "start_time": startTime, "end_time": endTime,
			"position_id": positionID, "user_id": userId}).Error; err != nil {
		return &models.Shifts{}, err
	}

	return shiftForUpdate, nil

}

func (p *PostgresDB) DeleteShift(ctx context.Context, id int64) error {

	tx := p.DB.Begin()
	if err := tx.WithContext(ctx).Model(&models.Shifts{}).Delete(&models.Shifts{}, id).Error; err != nil {
		return err
	} else {
		tx.Commit()
	}

	return nil
}

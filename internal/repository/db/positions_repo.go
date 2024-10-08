package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewPositionRepo() interfaces.IPositionsRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreatePosition(ctx context.Context, position *models.Positions) (*models.Positions, error) {
	if position == nil {
		return &models.Positions{}, errors.New("position is empty")
	}

	err := p.DB.WithContext(ctx).Create(&position).Error
	if err != nil {
		return &models.Positions{}, err
	}
	return position, nil
}

func (p *PostgresDB) AllPositions(ctx context.Context) ([]*models.Positions, error) {
	var positions []*models.Positions

	if err := p.DB.WithContext(ctx).Model(&models.Positions{}).Preload("Users").Preload("Shifts").Find(&positions).Error; err != nil {
		return []*models.Positions{}, err
	}

	return positions, nil
}

func (p *PostgresDB) GetPositionByID(ctx context.Context, id int64) (*models.Positions, error) {
	position := &models.Positions{}
	if err := p.DB.WithContext(ctx).Where("id = ?", id).Preload("Users").Preload("Shifts").Find(&position).Error; err != nil {
		return &models.Positions{}, err
	}

	return position, nil
}

func (p *PostgresDB) GetPositionByIdForEdit(ctx context.Context, id int64) (*models.Positions, error) {
	position := &models.Positions{}
	var usersArray []int64
	if err := p.DB.WithContext(ctx).Where("id = ?", id).Preload("Users").Preload("Shifts").Find(&position).Error; err != nil {
		return &models.Positions{}, err
	}

	for _, user := range position.Users {
		usersArray = append(usersArray, user.ID)
	}
	position.UsersArray = usersArray

	return position, nil
}

func (p *PostgresDB) AllPositionsForUserAddEdit(ctx context.Context) ([]*models.Positions, error) {
	var positions []*models.Positions
	if err := p.DB.WithContext(ctx).Model(&models.Positions{}).Select("id", "position_name").Find(&positions).Error; err != nil {
		return []*models.Positions{}, err
	}

	return positions, nil
}

func (p *PostgresDB) AllPositionsForDailySchedule(ctx context.Context) ([]*models.Positions, error) {
	var positions []*models.Positions
	if err := p.DB.WithContext(ctx).Model(&models.Positions{}).Select("position_name").Find(&positions).Error; err != nil {
		return []*models.Positions{}, err
	}
	return positions, nil
}

func (p *PostgresDB) UpdatePosition(ctx context.Context, id int64, idEdit int64, positionName string) (*models.Positions, error) {
	positionForUpdate, err := p.GetPositionByID(ctx, id)
	if err != nil {
		return &models.Positions{}, err
	}

	if err = p.DB.WithContext(ctx).Model(positionForUpdate).Where("id = ?", id).
		UpdateColumns(map[string]interface{}{"id": idEdit, "position_name": positionName}).Error; err != nil {
		return &models.Positions{}, err
	}

	return positionForUpdate, nil
}

func (p *PostgresDB) DeletePosition(ctx context.Context, id int64) error {
	tx := p.DB.WithContext(ctx).Begin()

	if err := tx.WithContext(ctx).Model(&models.Positions{}).Where("id = ?", id).Delete(&models.Positions{}).Error; err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

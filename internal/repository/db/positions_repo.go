package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"

	"github.com/google/uuid"
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

	if err := p.DB.WithContext(ctx).Model(&models.Positions{}).Preload("Users").Find(&positions).Error; err != nil {
		return []*models.Positions{}, err
	}

	return positions, nil
}

func (p *PostgresDB) GetPositionByID(ctx context.Context, id uuid.UUID) (*models.Positions, error) {
	position := &models.Positions{}
	if err := p.DB.WithContext(ctx).Where("id = ?", id).Preload("Users").Find(&position).Error; err != nil {
		return &models.Positions{}, err
	}

	return position, nil
}

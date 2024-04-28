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

func (p *PostgresDB) CreatePosition(ctx context.Context, position *models.Position) (*models.Position, error) {
	if position == nil {
		return &models.Position{}, errors.New("position is empty")
	}

	err := p.DB.WithContext(ctx).Create(&position).Error
	if err != nil {
		return &models.Position{}, err
	}
	return position, nil
}

package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IPositionsRepository interface {
	CreatePosition(ctx context.Context, position *models.Positions) (*models.Positions, error)
	AllPositions(ctx context.Context) ([]*models.Positions, error)
	GetPositionByID(ctx context.Context, id int64) (*models.Positions, error)
	GetPositionByIdForEdit(ctx context.Context, id int64) (*models.Positions, error)
	AllPositionsForUserAddEdit(ctx context.Context) ([]*models.Positions, error)
	UpdatePosition(ctx context.Context, id int64, idEdit int64, positionName string) (*models.Positions, error)
	DeletePosition(ctx context.Context, id int64) error
}

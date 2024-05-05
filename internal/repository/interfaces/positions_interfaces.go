package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"

	"github.com/google/uuid"
)

type IPositionsRepository interface {
	CreatePosition(ctx context.Context, position *models.Positions) (*models.Positions, error)
	AllPositions(ctx context.Context) ([]*models.Positions, error)
	GetPositionByID(ctx context.Context, id uuid.UUID) (*models.Positions, error)
}

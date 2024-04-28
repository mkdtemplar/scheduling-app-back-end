package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IPositionsRepository interface {
	CreatePosition(ctx context.Context, position *models.Positions) (*models.Positions, error)
}

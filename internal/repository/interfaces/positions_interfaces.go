package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IPositionsRepository interface {
	CreatePosition(ctx context.Context, position *models.Position) (*models.Position, error)
}

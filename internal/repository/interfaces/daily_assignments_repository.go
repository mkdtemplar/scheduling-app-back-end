package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IDailyAssignmentsRepository interface {
	GetUsersByPositionName(ctx context.Context, positionName string) ([]models.Users, error)
	CreateAssignment(ctx context.Context, a *models.DailyAssignment) error
}

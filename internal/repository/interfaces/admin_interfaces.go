package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IAdminRepository interface {
	CreateAdmin(ctx context.Context, admin *models.Administrator) (*models.Administrator, error)
}

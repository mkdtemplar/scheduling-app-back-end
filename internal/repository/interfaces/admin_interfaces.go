package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IAdminInterfaces interface {
	CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (*models.Admin, error)
}

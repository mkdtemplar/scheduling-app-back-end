package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IAdminInterfaces interface {
	CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (*models.Admin, error)
	UpdateAdmin(ctx context.Context, id int64, username string, password string) (*models.Admin, error)
	GetAdminById(ctx context.Context, id int64) (*models.Admin, error)
	DeleteAdmin(ctx context.Context, id int64) error
	AllAdmins(ctx context.Context) ([]*models.Admin, error)
}

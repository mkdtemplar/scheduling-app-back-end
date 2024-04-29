package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IUserRepository interface {
	CreateAdmin(ctx context.Context, admin *models.Users) (*models.Users, error)
}

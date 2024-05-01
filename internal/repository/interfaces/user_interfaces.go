package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, admin *models.Users) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}

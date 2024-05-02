package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, admin *models.Users) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*models.Users, error)
}

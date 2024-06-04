package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"
)

type IUserInterfaces interface {
	CreateUser(ctx context.Context, admin *models.Users) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserById(ctx context.Context, id int64) (*models.Users, error)
	AllUsers(ctx context.Context) ([]*models.Users, error)
	GetUserByIdForEdit(ctx context.Context, id int64) (*models.Users, error)
	UpdateUser(ctx context.Context, id int64, idUpdated int64, nameSurname string, email string,
		currentPosition string, positionId int64) (*models.Users, error)
	Delete(ctx context.Context, id int64) error
	GetUserIds(ctx context.Context) ([]*models.Users, error)
}

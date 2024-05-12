package interfaces

import (
	"context"
	"scheduling-app-back-end/internal/models"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, admin *models.Users) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserById(ctx context.Context, id int64) (*models.Users, error)
	AllUsers(ctx *gin.Context) ([]*models.Users, error)
	GetUserByIdForEdit(ctx context.Context, id int64) (*models.Users, error)
	UpdateUser(ctx context.Context, id int64, firstName string, lastName string, email string,
		currentPosition string, role string, positionId int64) (*models.Users, error)
}

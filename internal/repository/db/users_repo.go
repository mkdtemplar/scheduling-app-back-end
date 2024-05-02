package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewUserRepo() interfaces.IUserRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateUser(ctx context.Context, admin *models.Users) (*models.Users, error) {
	if admin == nil {
		return &models.Users{}, errors.New("admin details empty")
	}
	admin.ID = utils.GenerateID()

	err := p.DB.WithContext(ctx).Model(&models.Users{}).Create(&admin).Error
	if err != nil {
		return &models.Users{}, err
	}
	return admin, nil
}

func (p *PostgresDB) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := &models.Users{}

	err := p.DB.WithContext(ctx).Model(&models.Users{}).Where("email= ?", email).Take(&user).Error
	if err != nil {
		return &models.Users{}, err
	}

	userFind := &models.Users{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Password:        user.Password,
		CurrentPosition: user.CurrentPosition,
		Role:            user.Role,
		Shifts:          user.Shifts,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		PositionID:      user.PositionID,
	}
	return userFind, nil
}

func (p *PostgresDB) GetUserById(ctx context.Context, id uuid.UUID) (*models.Users, error) {
	user := &models.Users{}
	err := p.DB.WithContext(ctx).Model(&models.Users{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return &models.Users{}, err
	}
	userFind := &models.Users{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Password:        user.Password,
		CurrentPosition: user.CurrentPosition,
		Role:            user.Role,
		Shifts:          user.Shifts,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		PositionID:      user.PositionID,
	}

	return userFind, nil
}

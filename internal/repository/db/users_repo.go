package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"

	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewAdminRepo() interfaces.IUserRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateAdmin(ctx context.Context, admin *models.Users) (*models.Users, error) {
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

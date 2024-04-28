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

func NewAdminRepo() interfaces.IAdminRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateAdmin(ctx context.Context, admin *models.Administrator) (*models.Administrator, error) {
	if admin == nil {
		return &models.Administrator{}, errors.New("admin details empty")
	}
	admin.ID = utils.GenerateID()

	err := p.DB.WithContext(ctx).Model(&models.Administrator{}).Create(&admin).Error
	if err != nil {
		return &models.Administrator{}, err
	}
	return admin, nil
}

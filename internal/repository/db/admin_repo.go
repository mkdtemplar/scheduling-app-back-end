package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewAdminRepo() interfaces.IAdminInterfaces {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	if admin == nil {
		return &models.Admin{}, errors.New("user details empty")
	}

	err := p.DB.WithContext(ctx).Model(&models.Admin{}).Create(&admin).Error
	if err != nil {
		return &models.Admin{}, err
	}
	return admin, nil
}

func (p *PostgresDB) GetAdminByEmail(ctx context.Context, username string) (*models.Admin, error) {
	admin := &models.Admin{}

	err := p.DB.WithContext(ctx).Model(&models.Admin{}).Where("user_name= ?", username).Take(&admin).Error
	if err != nil {
		return &models.Admin{}, err
	}

	adminFind := &models.Admin{
		ID:       admin.ID,
		UserName: admin.UserName,
		Password: admin.Password,
	}
	return adminFind, nil
}

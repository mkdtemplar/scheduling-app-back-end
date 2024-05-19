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

func (p *PostgresDB) AllAdmins(ctx context.Context) ([]*models.Admin, error) {
	var admins []*models.Admin

	if err := p.DB.WithContext(ctx).Find(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
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

func (p *PostgresDB) GetAdminById(ctx context.Context, id int64) (*models.Admin, error) {
	admin := &models.Admin{}
	err := p.DB.WithContext(ctx).Model(&models.Admin{}).Where("id = ?", id).Take(&admin).Error
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

func (p *PostgresDB) UpdateAdmin(ctx context.Context, id int64, username string, password string) (*models.Admin, error) {
	var adminForUpdate = &models.Admin{}

	if err := p.DB.WithContext(ctx).Model(adminForUpdate).
		Where("id = ?", id).
		Updates(map[string]interface{}{"id": id, "user_name": username, "password": password}).Error; err != nil {
		return &models.Admin{}, err
	}
	return adminForUpdate, nil
}

func (p *PostgresDB) DeleteAdmin(ctx context.Context, id int64) error {

	tx := p.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.Admin{}).Delete(&models.Admin{}, id)

	if err := delTx.Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

package db

import (
	"context"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewDailyAssignmentsRepo() interfaces.IDailyAssignmentsRepository {
	return &PostgresDB{DB: GetDb()}
}
func (p *PostgresDB) GetUsersByPositionName(ctx context.Context, positionName string) ([]models.Users, error) {
	var users []models.Users
	err := p.DB.WithContext(ctx).
		Where("position_name = ?", positionName).
		Find(&users).Error
	return users, err
}

func (p *PostgresDB) CreateAssignment(ctx context.Context, a *models.DailyAssignment) error {
	if err := p.DB.WithContext(ctx).Model(&models.DailyAssignment{}).Create(&a).Error; err != nil {
		return err
	}
	return nil
}

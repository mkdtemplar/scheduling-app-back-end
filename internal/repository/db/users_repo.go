package db

import (
	"context"
	"errors"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewUserRepo() interfaces.IUserRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	if user == nil {
		return &models.Users{}, errors.New("user details empty")
	}

	err := p.DB.WithContext(ctx).Model(&models.Users{}).Create(&user).Error
	if err != nil {
		return &models.Users{}, err
	}
	return user, nil
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

func (p *PostgresDB) GetUserById(ctx context.Context, id int64) (*models.Users, error) {
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

func (p *PostgresDB) AllUsers(ctx *gin.Context) ([]*models.Users, error) {
	var users []*models.Users
	err := p.DB.WithContext(ctx).Model(&models.Users{}).Preload("Shifts").Find(&users).Error
	if err != nil {
		return []*models.Users{}, err
	}
	return users, nil
}

func (p *PostgresDB) GetUserByIdForEdit(ctx context.Context, id int64) (*models.Users, error) {
	user := &models.Users{}

	positions, err := p.AllPositions(ctx)
	if err != nil {
		return nil, err
	}
	err = p.DB.WithContext(ctx).Model(&models.Users{}).Where("id = ?", id).Take(&user).Error
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

	for _, position := range positions {
		userFind.UserPositionArray = append(userFind.UserPositionArray, position.ID)
	}
	return userFind, nil
}

func (p *PostgresDB) UpdateUser(ctx context.Context, id int64, firstName string, lastName string, email string,
	currentPosition string, role string, positionId int64) (*models.Users, error) {

	var userForUpdate = &models.Users{}

	if err := p.DB.WithContext(ctx).Model(userForUpdate).Where("id = ?", id).
		Updates(map[string]interface{}{"first_name": firstName, "last_name": lastName, "email": email,
			"current_position": currentPosition, "role": role, "position_id": positionId}).Error; err != nil {
		return &models.Users{}, err
	}

	return userForUpdate, nil
}

func (p *PostgresDB) Delete(ctx context.Context, id int64) error {
	var err error

	tx := p.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.Users{}).Delete(&models.Users{}, id)

	if err = delTx.Error; err != nil {
		return err
	} else {
		tx.Commit()
	}

	return nil
}

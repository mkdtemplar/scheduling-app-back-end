package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserHandler(IAdminRepository interfaces.IUserRepository) *UserHandler {
	return &UserHandler{
		IUserRepository: IAdminRepository,
	}
}

type createUserRequest struct {
	ID              uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName       string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName        string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email           string      `gorm:"type:text" json:"email" binding:"required,email"`
	Password        string      `gorm:"type:text" json:"password" binding:"required"`
	CurrentPosition string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role            models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	CreatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID      uuid.UUID   `gorm:"type:uuid" json:"position_id"`
}

type createUserResponse struct {
	ID                uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName         string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName          string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email             string      `gorm:"type:text" json:"email" binding:"required,email"`
	CurrentPosition   string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role              models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID        uuid.UUID   `gorm:"type:uuid" json:"position_id"`
}

func newUserResponse(user *models.Users) *createUserResponse {
	return &createUserResponse{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		CurrentPosition:   user.CurrentPosition,
		Role:              user.Role,
		PositionID:        user.PositionID,
		PasswordChangedAt: time.Now().UTC(),
		CreatedAt:         time.Now().UTC(),
	}
}

func (admin *UserHandler) Create(ctx *gin.Context) {
	var req *createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := &models.Users{
		ID:              utils.GenerateID(),
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		Password:        hashedPassword,
		CurrentPosition: req.CurrentPosition,
		Role:            req.Role,
		PositionID:      req.PositionID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	newUser, err := admin.IUserRepository.CreateAdmin(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(newUser)
	ctx.JSON(http.StatusCreated, response)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

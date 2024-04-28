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

func NewAdminHandler(IAdminRepository interfaces.IAdminRepository) *AdminHandler {
	return &AdminHandler{
		IAdminRepository: IAdminRepository,
	}
}

type createAdminRequest struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Username string    `json:"username" gorm:"varchar(25);unique" binding:"required,email"`
	Password string    `json:"password" gorm:"varchar(32)" binding:"required,min=8,max=32"`
}

type createAdminResponse struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Username          string    `json:"username" gorm:"varchar(25);unique" binding:"required,email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newAdminResponse(admin *models.Administrator) *createAdminResponse {
	return &createAdminResponse{
		ID:                admin.ID,
		Username:          admin.Username,
		PasswordChangedAt: time.Now().UTC(),
		CreatedAt:         time.Now().UTC(),
	}
}

func (admin *AdminHandler) Create(ctx *gin.Context) {
	var req *createAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := &models.Administrator{
		ID:       utils.GenerateID(),
		Username: req.Username,
		Password: hashedPassword,
	}

	newAdmin, err := admin.IAdminRepository.CreateAdmin(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newAdminResponse(newAdmin)
	ctx.JSON(http.StatusCreated, response)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

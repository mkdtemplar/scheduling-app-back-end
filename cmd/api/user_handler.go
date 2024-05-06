package api

import (
	"errors"
	"net/http"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserHandler(IUserRepository interfaces.IUserRepository, IJWTInterfaces middleware.IJWTInterfaces) *UserHandler {
	return &UserHandler{
		IUserRepository: IUserRepository,
		IJWTInterfaces:  IJWTInterfaces,
	}
}

type createUserRequest struct {
	ID              int64       `gorm:"type:bigint;primaryKey" json:"id"`
	FirstName       string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName        string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email           string      `gorm:"type:text" json:"email" binding:"required,email"`
	Password        string      `gorm:"type:text" json:"password" binding:"required"`
	CurrentPosition string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role            models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	CreatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID      int64       `gorm:"type:bigint" json:"position_id"`
}

type CreateUserResponse struct {
	ID                int64       `gorm:"type:bigint;primaryKey" json:"id"`
	FirstName         string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName          string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email             string      `gorm:"type:text" json:"email" binding:"required,email"`
	CurrentPosition   string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role              models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	PasswordChangedAt time.Time   `json:"password_changed_at,omitempty"`
	CreatedAt         time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID        int64       `gorm:"type:bigint" json:"position_id"`
}

func NewUserResponse(user *models.Users) *CreateUserResponse {
	return &CreateUserResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		CurrentPosition: user.CurrentPosition,
		Role:            user.Role,
		PositionID:      user.PositionID,
	}
}

func (user *UserHandler) Create(ctx *gin.Context) {
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

	newUser, err := user.IUserRepository.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := NewUserResponse(newUser)
	ctx.JSON(http.StatusCreated, response)
}

func (user *UserHandler) AllUsers(ctx *gin.Context) {

	var allUsers []*CreateUserResponse

	users, err := user.IUserRepository.AllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, u := range users {
		i := NewUserResponse(u)
		allUsers = append(allUsers, i)
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (user *UserHandler) Authorization(ctx *gin.Context) {

	var requestPayload struct {
		Email    string `json:"email" binding:"required" gorm:"type:email"`
		Password string `json:"password" binding:"required" gorm:"type:password"`
	}

	if err := ctx.ShouldBindJSON(&requestPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userFromDb, err := user.IUserRepository.GetUserByEmail(ctx, requestPayload.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	valid, err := utils.CheckPassword(requestPayload.Password, userFromDb.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	testUser := middleware.JwtUser{
		ID:        userFromDb.ID,
		FirstName: userFromDb.FirstName,
		LastName:  userFromDb.LastName,
		Email:     userFromDb.Email,
	}

	tokens, err := user.IJWTInterfaces.GenerateTokenPairs(&testUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credentials")))
		return
	}

	user.IJWTInterfaces.GetRefreshCookie(tokens.RefreshToken, ctx)

	ctx.JSON(http.StatusAccepted, tokens)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

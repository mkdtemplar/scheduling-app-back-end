package api

import (
	"net/http"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(IUserRepository interfaces.IUserRepository, IJWTInterfaces middleware.IJWTInterfaces) *UserHandler {
	return &UserHandler{
		IUserRepository: IUserRepository,
		IJWTInterfaces:  IJWTInterfaces,
	}
}

type createUserRequest struct {
	ID              int64       `gorm:"type:bigint;primaryKey" json:"id,string" binding:"required"`
	FirstName       string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName        string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email           string      `gorm:"type:text" json:"email" binding:"required,email"`
	Password        string      `gorm:"type:text" json:"password" binding:"required"`
	CurrentPosition string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role            models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	CreatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID      int64       `gorm:"type:bigint" json:"position_id,string" binding:"required"`
}

type CreateUserResponse struct {
	ID                int64       `gorm:"type:bigint;primaryKey" json:"id,string"`
	FirstName         string      `gorm:"type:text" json:"first_name" binding:"required"`
	LastName          string      `gorm:"type:text" json:"last_name" binding:"required"`
	Email             string      `gorm:"type:text" json:"email" binding:"required,email"`
	CurrentPosition   string      `gorm:"type:text" json:"current_position" binding:"required"`
	Role              models.Role `sql:"type:user_role" db:"role" json:"role" binding:"required"`
	PasswordChangedAt time.Time   `json:"password_changed_at,omitempty"`
	CreatedAt         time.Time   `gorm:"type:timestamp" json:"-"`
	PositionID        int64       `gorm:"type:bigint" json:"position_id,string"`
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

func (usr *UserHandler) Create(ctx *gin.Context) {
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
		ID:              req.ID,
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

	newUser, err := usr.IUserRepository.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := NewUserResponse(newUser)
	ctx.JSON(http.StatusCreated, response)
}

func (usr *UserHandler) AllUsers(ctx *gin.Context) {

	var allUsers []*CreateUserResponse

	users, err := usr.IUserRepository.AllUsers(ctx)
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

func (usr *UserHandler) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userById, err := usr.IUserRepository.GetUserById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	response := NewUserResponse(userById)

	ctx.JSON(http.StatusOK, response)
}

func (usr *UserHandler) GetUserByIdForEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userForEdit, err := usr.IUserRepository.GetUserByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := NewUserResponse(userForEdit)

	ctx.JSON(http.StatusOK, response)
}

func (usr *UserHandler) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userFromDb, err := usr.IUserRepository.GetUserByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	userForEdit, err := utils.ParseUserPrefRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userFromDb, err = usr.IUserRepository.UpdateUser(ctx, int64(id), userForEdit.FirstName, userForEdit.LastName,
		userForEdit.Email, userForEdit.CurrentPosition, string(userForEdit.Role), userForEdit.PositionID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := NewUserResponse(userFromDb)

	ctx.JSON(http.StatusOK, response)

}

func (usr *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = usr.IUserRepository.Delete(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

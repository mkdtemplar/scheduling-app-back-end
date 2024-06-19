package api

import (
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(IUserRepository interfaces.IUserInterfaces) *UserHandler {
	return &UserHandler{
		IUserInterfaces: IUserRepository,
	}
}

func (usr *UserHandler) Create(ctx *gin.Context) {
	var req *dto.CreateUserRequest
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
		ID:           req.ID,
		NameSurname:  req.NameSurname,
		Email:        req.Email,
		Password:     hashedPassword,
		PositionName: req.PositionName,
		UserID:       req.UserID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	newUser, err := usr.IUserInterfaces.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewUserResponse(newUser)
	ctx.JSON(http.StatusCreated, response)
}

func (usr *UserHandler) AllUsers(ctx *gin.Context) {

	var allUsers []*dto.CreateUserResponse

	users, err := usr.IUserInterfaces.AllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, u := range users {
		i := dto.NewUserResponse(u)
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

	userById, err := usr.IUserInterfaces.GetUserById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	response := dto.NewUserResponse(userById)

	ctx.JSON(http.StatusOK, response)
}

func (usr *UserHandler) GetUserByIdForEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userForEdit, err := usr.IUserInterfaces.GetUserByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewUserResponse(userForEdit)

	ctx.JSON(http.StatusOK, response)
}

func (usr *UserHandler) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userFromDb, err := usr.IUserInterfaces.GetUserByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	userForEdit, err := utils.ParseUserPrefRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userUpdated, err := usr.IUserInterfaces.UpdateUser(ctx, userFromDb.ID, userForEdit.ID, userForEdit.NameSurname,
		userForEdit.Email, userForEdit.PositionName, userForEdit.UserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewUserResponse(userUpdated)

	ctx.JSON(http.StatusOK, response)

}

func (usr *UserHandler) GetUserIds(ctx *gin.Context) {
	var allIds []*dto.UserResponseForShift
	ids, err := usr.IUserInterfaces.GetUserIds(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, id := range ids {
		allIds = append(allIds, dto.NewUserResponseForShift(id))
	}

	ctx.JSON(http.StatusOK, allIds)
}

func (usr *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = usr.IUserInterfaces.Delete(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"error": false, "message": "user deleted"})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

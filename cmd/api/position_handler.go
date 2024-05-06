package api

import (
	"fmt"
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/repository/interfaces"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createPositionRequest struct {
	PositionName string    `json:"position_name" gorm:"type:text"`
	StartTime    string    `json:"start_time" gorm:"type:varchar(10)"`
	UpdatedAt    time.Time `json:"-" gorm:"type:timestamp"`
}

type getPositionRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type positionResponse struct {
	ID           int64                 `gorm:"type:bigint;primaryKey" json:"id"`
	PositionName string                `json:"position_name" gorm:"type:text"`
	Users        []*CreateUserResponse `gorm:"foreignKey:PositionID;references:ID" json:"users,omitempty"`
	Shifts       []*models.Shifts      `gorm:"foreignKey:PositionID;references:ID" json:"shifts,omitempty"`
	CreatedAt    time.Time             `json:"-" gorm:"type:timestamp"`
	UpdatedAt    time.Time             `json:"-" gorm:"type:timestamp"`
	UsersArray   []int64               `gorm:"-" json:"users_array,omitempty"`
}

func newPositionResponse(positions *models.Positions) *positionResponse {
	var allUsers []*CreateUserResponse
	for _, u := range positions.Users {
		i := NewUserResponse(u)
		allUsers = append(allUsers, i)
	}

	response := &positionResponse{
		ID:           positions.ID,
		PositionName: positions.PositionName,
		Users:        allUsers,
		Shifts:       positions.Shifts,
		UsersArray:   positions.UsersArray,
	}

	return response
}

func NewPositionHandler(IPositionRepository interfaces.IPositionsRepository) *PositionHandler {
	return &PositionHandler{
		IPositionsRepository: IPositionRepository,
	}
}

func (i *PositionHandler) CreatePosition(ctx *gin.Context) {
	var req *createPositionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &models.Positions{
		PositionName: req.PositionName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	newPosition, err := i.IPositionsRepository.CreatePosition(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newPosition)
}

func (i *PositionHandler) AllPositions(ctx *gin.Context) {
	var allPositions []*positionResponse
	positions, err := i.IPositionsRepository.AllPositions(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, position := range positions {
		position.UsersArray = append(position.UsersArray, position.ID)
	}
	for _, position := range positions {
		allPositions = append(allPositions, newPositionResponse(position))
	}

	ctx.JSON(http.StatusOK, allPositions)
}

func (i *PositionHandler) GetPositionById(ctx *gin.Context) {
	var req getPositionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req.ID)

	position, err := i.IPositionsRepository.GetPositionByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

func (i *PositionHandler) GetPositionByIdForEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	position, err := i.IPositionsRepository.GetPositionByIdForEdit(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNoContent, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

package api

import (
	"net/http"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewAdminHandler(IAdminInterfaces interfaces.IAdminInterfaces, IJWTInterfaces middleware.IJWTInterfaces) *AdminHandler {
	return &AdminHandler{IAdminInterfaces: IAdminInterfaces, IJWTInterfaces: IJWTInterfaces}
}

func (adm *AdminHandler) AllAdmins(ctx *gin.Context) {
	var allAdmins []*dto.CreateAdminResponse

	admins, err := adm.IAdminInterfaces.AllAdmins(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, admin := range admins {
		a := dto.NewAdminResponse(admin)
		allAdmins = append(allAdmins, a)
	}
	ctx.JSON(http.StatusOK, allAdmins)
}

func (adm *AdminHandler) CreateAdmin(ctx *gin.Context) {
	var req *dto.CreateAdminRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := &models.Admin{
		ID:       req.ID,
		UserName: req.UserName,
		Password: hashedPassword,
	}

	newAdmin, err := adm.IAdminInterfaces.CreateAdmin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewAdminResponse(newAdmin)

	ctx.JSON(http.StatusCreated, response)
}

func (adm *AdminHandler) UpdateAdmin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	adminFromDb, err := adm.IAdminInterfaces.GetAdminById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	adminForEdit, err := utils.ParseAdminRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(adminForEdit.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	adminFromDb, err = adm.IAdminInterfaces.UpdateAdmin(ctx, adminForEdit.ID, adminForEdit.UserName, hashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := dto.NewAdminResponse(adminFromDb)
	ctx.JSON(http.StatusOK, response)
}

func (adm *AdminHandler) DeleteAdmin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := adm.IAdminInterfaces.DeleteAdmin(ctx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"error": false, "message": "admin deleted"})
}

func (adm *AdminHandler) GetAdminById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := adm.IAdminInterfaces.GetAdminById(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

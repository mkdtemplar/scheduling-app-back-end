package api

import (
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/interfaces"
)

func NewAdminHandler(IAdminInterfaces interfaces.IAdminInterfaces, IJWTInterfaces middleware.IJWTInterfaces) *AdminHandler {
	return &AdminHandler{IAdminInterfaces: IAdminInterfaces, IJWTInterfaces: IJWTInterfaces}
}

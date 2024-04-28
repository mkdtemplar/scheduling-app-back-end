package api

import "scheduling-app-back-end/internal/repository/db"

func (admin *AdminHandler) AdminHandlerConstructor() *AdminHandler {
	adminRepo := db.NewAdminRepo()

	return NewAdminHandler(adminRepo)
}

func (i *PositionHandler) PositionHandlerConstructor() *PositionHandler {
	positionRepo := db.NewPositionRepo()
	return NewPositionHandler(positionRepo)
}

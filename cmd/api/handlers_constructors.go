package api

import "scheduling-app-back-end/internal/repository/db"

func (admin *UserHandler) AdminHandlerConstructor() *UserHandler {
	adminRepo := db.NewAdminRepo()

	return NewUserHandler(adminRepo)
}

func (i *PositionHandler) PositionHandlerConstructor() *PositionHandler {
	positionRepo := db.NewPositionRepo()
	return NewPositionHandler(positionRepo)
}

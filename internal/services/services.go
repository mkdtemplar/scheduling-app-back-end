package services

import (
	"log"
	"scheduling-app-back-end/api"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/utils"
)

type AdminHandler api.AdminHandler
type PositionHandler api.PositionHandler
type UserHandler api.UserHandler

func (adm *AdminHandler) AdminHandlerConstructor() *api.AdminHandler {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	adminRepo := db.NewAdminRepo()

	tokenPairs := middleware.NewAuthorization(config.Issuer, config.Audience, config.JWTSecret, config.TokenExpiry,
		config.RefreshExpiry, config.CookieDomain, config.CookiePath, config.CookieName)

	return api.NewAdminHandler(adminRepo, tokenPairs)
}

func (i *PositionHandler) PositionHandlerConstructor() *api.PositionHandler {
	positionRepo := db.NewPositionRepo()
	return api.NewPositionHandler(positionRepo)
}

func (usr *UserHandler) UserHandlerConstructor() *api.UserHandler {
	userRepo := db.NewUserRepo()
	return api.NewUserHandler(userRepo)
}

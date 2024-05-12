package api

import (
	"log"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/utils"
)

func (usr *UserHandler) UserHandlerConstructor() *UserHandler {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	userRepo := db.NewUserRepo()

	tokenPairs := middleware.NewAuthorization(config.Issuer, config.Audience, config.JWTSecret, config.TokenExpiry,
		config.RefreshExpiry, config.CookieDomain, config.CookiePath, config.CookieName)

	return NewUserHandler(userRepo, tokenPairs)
}

func (i *PositionHandler) PositionHandlerConstructor() *PositionHandler {
	positionRepo := db.NewPositionRepo()
	return NewPositionHandler(positionRepo)
}

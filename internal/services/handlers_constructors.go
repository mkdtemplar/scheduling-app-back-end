package services

import (
	"log"
	"scheduling-app-back-end/api"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/utils"
)

type Usr api.UserHandler
type Pos api.PositionHandler

func (usr *Usr) UserHandlerConstructor() *api.UserHandler {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	userRepo := db.NewUserRepo()

	tokenPairs := middleware.NewAuthorization(config.Issuer, config.Audience, config.JWTSecret, config.TokenExpiry,
		config.RefreshExpiry, config.CookieDomain, config.CookiePath, config.CookieName)

	return api.NewUserHandler(userRepo, tokenPairs)
}

func (i *Pos) PositionHandlerConstructor() *api.PositionHandler {
	positionRepo := db.NewPositionRepo()
	return api.NewPositionHandler(positionRepo)
}

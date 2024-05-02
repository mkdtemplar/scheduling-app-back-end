package api

import (
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/repository/interfaces"
)

type Handler struct {
	DB db.PostgresDB
}

type UserHandler struct {
	Handler
	interfaces.IUserRepository
	middleware.IJWTInterfaces
	middleware.JwtUser
}

type PositionHandler struct {
	Handler
	interfaces.IPositionsRepository
}

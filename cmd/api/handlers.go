package api

import (
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/repository/interfaces"
)

type Handler struct {
	DB db.PostgresDB
}

type AdminHandler struct {
	Handler
	interfaces.IAdminRepository
}

type PositionHandler struct {
	Handler
	interfaces.IPositionsRepository
}

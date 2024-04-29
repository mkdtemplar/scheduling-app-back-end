package server

import (
	"scheduling-app-back-end/cmd/api"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/repository/db"
	"scheduling-app-back-end/internal/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	Router *gin.Engine
}

func NewServer(config utils.Config) (*Server, error) {
	server := &Server{config: config}

	server.setupRouter()
	return server, nil
}

func (server *Server) InitServer() {
	db.ConnectToPostgres()

	server.setupRouter()
}

func (server *Server) setupRouter() {
	var adminHandlers api.UserHandler
	var positionsHandlers api.PositionHandler
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	adminHandler := adminHandlers.AdminHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()
	router.GET("/")
	router.POST("/admin/create", adminHandler.Create)
	router.POST("/admin/position/create", positionHandler.CreatePosition)
	router.GET("/admin/position/all-positions", positionHandler.AllPositions)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

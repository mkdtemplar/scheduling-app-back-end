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

var AdminHandlers api.AdminHandler
var PositionsHandlers api.PositionHandler

func (server *Server) InitServer() {
	db.ConnectToPostgres()

	server.setupRouter()
}

func (server *Server) setupRouter() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	adminHandler := AdminHandlers.AdminHandlerConstructor()
	positionHandler := PositionsHandlers.PositionHandlerConstructor()
	router.GET("/")
	router.POST("/admin/create", adminHandler.Create)
	router.POST("/admin/position/create", positionHandler.CreatePosition)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

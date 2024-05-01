package server

import (
	"scheduling-app-back-end/cmd/api"
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	Router *gin.Engine
	auth   middleware.IJWTInterfaces
}

func NewServer(config utils.Config) (*Server, error) {

	server := &Server{config: config}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	var userHandlers api.UserHandler
	var positionsHandlers api.PositionHandler
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	corsRoutes := router.Group("/", middleware.CORSMiddleware())

	userHandler := userHandlers.AdminHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()

	corsRoutes.POST("/authenticate", userHandler.Authorization)
	corsRoutes.GET("/")
	corsRoutes.POST("/admin/create", userHandler.Create)
	corsRoutes.POST("/admin/position/create", positionHandler.CreatePosition)
	corsRoutes.GET("/admin/position/all-positions", positionHandler.AllPositions)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

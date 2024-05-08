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

	userHandler := userHandlers.UserHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()

	router.Use(middleware.CORSMiddleware())

	router.POST("/authenticate", userHandler.Authorization)
	router.GET("/refresh", userHandler.RefreshToken)
	router.GET("/logout", userHandler.Logout)
	router.GET("/")
	router.GET("/positions", positionHandler.AllPositions)
	router.GET("/position/:id", positionHandler.GetPositionById)
	router.GET("/all-users", userHandler.AllUsers)

	authRoutes := router.Group("/admin").Use(userHandler.IJWTInterfaces.AuthRequired())
	authRoutes.POST("/add-user", userHandler.Create)
	authRoutes.POST("/add-position", positionHandler.CreatePosition)
	authRoutes.GET("/edit-position/:id", positionHandler.GetPositionByIdForEdit)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

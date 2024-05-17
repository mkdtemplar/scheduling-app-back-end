package server

import (
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/services"
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
	var userHandlers services.UserHandler
	var positionsHandlers services.PositionHandler
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	userHandler := userHandlers.UserHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()

	router.Use(middleware.CORSMiddleware())

	router.GET("/")
	router.POST("/authenticate", userHandler.Authorization)
	router.GET("/refresh", userHandler.RefreshToken)
	router.GET("/logout", userHandler.Logout)
	router.GET("/positions", positionHandler.AllPositions)
	router.GET("/position/:id", positionHandler.GetPositionById)
	router.GET("/position-for-user", positionHandler.AllPositionsForUserAddEdit)
	router.GET("/all-users", userHandler.AllUsers)
	router.GET("/user/:id", userHandler.GetUserById)

	authRoutes := router.Group("/admin").Use(userHandler.IJWTInterfaces.AuthRequired())
	authRoutes.PUT("/add-user", userHandler.Create)
	authRoutes.PUT("/add-position/0", positionHandler.CreatePosition)
	authRoutes.GET("/edit-position/:id", positionHandler.GetPositionByIdForEdit)
	authRoutes.GET("/user-edit/:id", userHandler.GetUserByIdForEdit)
	authRoutes.PATCH("/edit-user/:id", userHandler.UpdateUser)

	authRoutes.DELETE("/delete-user/:id", userHandler.DeleteUser)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

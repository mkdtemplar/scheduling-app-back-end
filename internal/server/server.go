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
	var adminHandlers services.AdminHandler
	var positionsHandlers services.PositionHandler
	var userHandlers services.UserHandler
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	gin.ForceConsoleColor()

	adminHandler := adminHandlers.AdminHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()
	userHandler := userHandlers.UserHandlerConstructor()

	router.Use(middleware.CORSMiddleware())

	router.GET("/")
	router.POST("/authenticate", adminHandler.Authorization)
	router.GET("/refresh", adminHandler.RefreshToken)
	router.GET("/logout", adminHandler.Logout)
	router.GET("/positions", positionHandler.AllPositions)
	router.GET("/position/:id", positionHandler.GetPositionById)
	router.GET("/position-for-user", positionHandler.AllPositionsForUserAddEdit)
	router.GET("/all-users", userHandler.AllUsers)
	router.GET("/user/:id", userHandler.GetUserById)

	authRoutes := router.Group("/admin").Use(adminHandler.IJWTInterfaces.AuthRequired())
	authRoutes.PUT("/add-user", userHandler.Create)
	authRoutes.PUT("/add-position/0", positionHandler.CreatePosition)
	authRoutes.GET("/edit-position/:id", positionHandler.GetPositionByIdForEdit)
	authRoutes.GET("/user-edit/:id", userHandler.GetUserByIdForEdit)
	authRoutes.PATCH("/edit-user/:id", userHandler.UpdateUser)

	authRoutes.DELETE("/delete-user/:id", userHandler.DeleteUser)
	authRoutes.GET("/all-admins", adminHandler.AllAdmins)
	authRoutes.PUT("/create-admin", adminHandler.CreateAdmin)
	authRoutes.GET("/get-admin/:id", adminHandler.GetAdminById)
	authRoutes.PATCH("/update-admin/:id", adminHandler.UpdateAdmin)
	authRoutes.DELETE("/delete-admin/:id", adminHandler.DeleteAdmin)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

package server

import (
	"scheduling-app-back-end/internal/middleware"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/services"
	"scheduling-app-back-end/internal/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config   utils.Config
	Router   *gin.Engine
	auth     middleware.IJWTInterfaces
	MailChan chan models.MailData
}

func NewServer(config utils.Config, mailChan chan models.MailData) (*Server, error) {

	server := &Server{config: config, MailChan: mailChan}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	var adminHandlers services.AdminHandler
	var positionsHandlers services.PositionHandler
	var userHandlers services.UserHandler
	var shiftsHandlers services.ShiftHandler
	var annualLeaveHandlers services.AnnualLeaveHandler
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	gin.ForceConsoleColor()

	adminHandler := adminHandlers.AdminHandlerConstructor()
	positionHandler := positionsHandlers.PositionHandlerConstructor()
	userHandler := userHandlers.UserHandlerConstructor()
	shiftHandler := shiftsHandlers.ShiftsHandlerConstructor()
	annualLeaveHandler := annualLeaveHandlers.AnnualLeaveConstructor()

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
	router.GET("/all-admins", adminHandler.AllAdmins)
	router.GET("/all-shifts", shiftHandler.GetAllShifts)
	router.GET("/get-admin/:id", adminHandler.GetAdminById)
	router.GET("/get-shift/:id", shiftHandler.GetShiftById)
	router.GET("/get-shift-name/:name", shiftHandler.GetShiftByName)
	router.GET("/user-ids", userHandler.GetUserIds)
	router.PUT("/create-annual-leave", annualLeaveHandler.CreateAnnualLeave)

	authRoutes := router.Group("/admin").Use(adminHandler.IJWTInterfaces.AuthRequired())
	authRoutes.PUT("/add-user", userHandler.Create)
	authRoutes.PUT("/add-position", positionHandler.CreatePosition)
	authRoutes.GET("/edit-position/:id", positionHandler.GetPositionByIdForEdit)
	authRoutes.PATCH("/update-position/:id", positionHandler.UpdatePosition)
	authRoutes.DELETE("/delete-position/:id", positionHandler.DeletePosition)
	authRoutes.GET("/user-edit/:id", userHandler.GetUserByIdForEdit)
	authRoutes.PATCH("/edit-user/:id", userHandler.UpdateUser)

	authRoutes.DELETE("/delete-user/:id", userHandler.DeleteUser)

	authRoutes.PUT("/create-admin", adminHandler.CreateAdmin)

	authRoutes.PATCH("/update-admin/:id", adminHandler.UpdateAdmin)
	authRoutes.DELETE("/delete-admin/:id", adminHandler.DeleteAdmin)

	authRoutes.PUT("/create-shift", shiftHandler.CreateShift)
	authRoutes.PATCH("/update-shift/:id", shiftHandler.UpdateShift)
	authRoutes.DELETE("/delete-shift/:id", shiftHandler.DeleteShift)

	server.Router = router
}

func (server *Server) Run(address string) error {
	return server.Router.Run(address)
}

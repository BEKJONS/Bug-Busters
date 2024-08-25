package api

import (
	_ "bug_busters/api/docs"
	"bug_busters/api/handler"
	"bug_busters/internal/service"
	"bug_busters/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authentication service
// @version 1.0
// @description Server for signIn or signUp
// @BasePath /
// @schemes http
func NewRouter(s service.AuthService, i service.IIService, u service.UserService, serv service.IService) *gin.Engine {
	r := gin.New()
	h := handler.NewHandler(logger.NewLogger(), s, i, serv, u)

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/upload", h.AddImage)

	// Authentication routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/add_license", h.AddLicense)
	}

	// Fines routes
	fines := r.Group("/fines")
	{
		fines.POST("", h.CreateFines)
		fines.PUT("/:id/accept", h.AcceptFinesById) // Assuming you need to accept fines by ID
		fines.GET("/paid", h.GetPaidFines)
		fines.GET("/unpaid", h.GetUnpaidFines)
		fines.GET("", h.GetAllFines) // Get all fines with optional pagination
	}
	// Service routes
	service := r.Group("/service")
	{
		service.GET("", h.GetAllServices)
		service.GET("/:id", h.GetService)
		service.POST("/create", h.CreateService)
		service.PUT("/update", h.UpdateService)
		service.DELETE("/delete/:id", h.DeleteService)
	}

	return r
}

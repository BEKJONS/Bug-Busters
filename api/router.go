package api

import (
	_ "bug_busters/api/docs"
	"bug_busters/api/handler"
	"bug_busters/api/middleware"
	"bug_busters/internal/service"
	"bug_busters/pkg/logger"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authentication service
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @description Server for signIn or signUp
// @in header
// @name Authorization
// @schemes http
// @BasePath /
func NewRouter(s service.AuthService, i service.IIService, u service.UserService, serv service.IService, enf *casbin.Enforcer) *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	h := handler.NewHandler(logger.NewLogger(), s, i, serv, u)

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authentication routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/add_license", h.AddLicense)
	}
	r.Use(middleware.PermissionMiddleware(enf))
	// Fines routes
	fines := r.Group("/fines")
	{
		fines.POST("", h.CreateFines)
		fines.PUT("/:id/accept", h.AcceptFinesById) // Assuming you need to accept fines by ID
		fines.GET("/paid", h.GetPaidFines)
		fines.GET("/unpaid", h.GetUnpaidFines)
		fines.GET("", h.GetAllFines) // Get all fines with optional pagination
		fines.POST("/send_acceptation", h.SendAcceptation)
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

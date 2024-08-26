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
	r.POST("/upload", h.AddImage)
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
	services := r.Group("/service")
	{
		services.GET("", h.GetAllServices)
		services.GET("/:id", h.GetService)
		services.POST("/create", h.CreateService)
		services.PUT("/update", h.UpdateService)
		services.DELETE("/delete/:id", h.DeleteService)
	}
	user := r.Group("/user")
	{
		user.GET("/profile", h.GetProfile)
		user.GET("/paid_fines", h.GetPaidFinesU)
		user.GET("/unpaid_fines", h.GetUnpaid)
	}
	admin := r.Group("/admin")
	{
		admin.GET("/profile/:id", h.GetProfileAdmin)
		admin.GET("/paid_fines/:id", h.GetPaidFinesAdmin)
		admin.GET("/unpaid_fines/:id", h.GetUnpaidAdmin)
		admin.DELETE("/:id", h.DeleteUser)

	}
	return r
}

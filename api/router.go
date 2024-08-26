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
func NewRouter(s service.AuthService, i service.IIService, u service.UserService, serv service.IService, sw service.SWService, enf *casbin.Enforcer) *gin.Engine {
	r := gin.New()

	h := handler.NewHandler(logger.NewLogger(), s, i, serv, u, sw)

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

	r.POST("/image", h.AddImage)
	r.GET("/image", h.GetImage)

	// Fines routes
	fines := r.Group("/fines")
	{
		fines.POST("", h.CreateFines)
		fines.PUT("/:id/accept", h.AcceptFinesById)
		fines.GET("/paid", h.GetPaidFines)
		fines.GET("/unpaid", h.GetUnpaidFines)
		fines.GET("", h.GetAllFines)
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

	// User routes
	user := r.Group("/user")
	{
		user.GET("/profile", h.GetProfile)
		user.GET("/paid_fines", h.GetPaidFinesU)
		user.GET("/unpaid_fines", h.GetUnpaid)
	}

	// Admin routes
	admin := r.Group("/admin")
	{
		admin.GET("/profile/:id", h.GetProfileAdmin)
		admin.GET("/paid_fines/:id", h.GetPaidFinesAdmin)
		admin.GET("/unpaid_fines/:id", h.GetUnpaidAdmin)
		admin.DELETE("/:id", h.DeleteUser)
	}

	// Single Window Service (SWS) routes
	sws := r.Group("/single_window")
	{
		sws.POST("/license", h.CreateLicense)
		sws.POST("/passport", h.CreatePassport)
		sws.GET("/licenses", h.GetAllLicenses)
		sws.GET("/passports", h.GetAllPassports)
		sws.DELETE("/license", h.DeleteLicense)
		sws.DELETE("/passport", h.DeletePassport)
	}

	return r
}

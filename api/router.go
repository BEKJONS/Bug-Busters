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
func NewRouter(s service.AuthService, i service.IIService) *gin.Engine {
	r := gin.New()
	h := handler.NewHandler(logger.NewLogger(), s, i)

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authentication routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
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
	// User routes
	user := r.Group("/user")
	{
		user.GET("/profile/:id", h.GetProfile)          // Get user profile by ID
		user.POST("/image", h.AddImage)                 // Add car image
		user.GET("/image/:id", h.GetImage)              // Get car image by user ID
		user.GET("/paid_fines/:id", h.GetPaidFines)     // Get paid fines by user ID
		user.GET("/unpaid_fines/:id", h.GetUnpaidFines) // Get unpaid fines by user ID
		user.DELETE("/:id", h.DeleteUser)               // Delete user by ID
	}

	return r
}

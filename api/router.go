package api

import (
	"bug_busters/api/handler"
	"bug_busters/internal/service"
	"bug_busters/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authenfication service
// @version 1.0
// @description server for siginIn or signUp
// @BasePath /auth
// @schemes http
func NewRouter(s service.AuthService) *gin.Engine {
	r := gin.New()
	h := handler.NewAuthHandler(logger.NewLogger(), s)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	return r
}

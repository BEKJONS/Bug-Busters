package handler

import (
	"bug_busters/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	srv service.AuthService
	log *slog.Logger
}

func NewAuthHandler(log *slog.Logger, sr service.AuthService) AuthHandler {
	return &authHandler{log: log, srv: sr}
}

package handler

import (
	"bug_busters/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handlers interface {
	Register(c *gin.Context)
	Login(c *gin.Context)

	CreateFines(c *gin.Context)
	AcceptFinesById(c *gin.Context)
	GetPaidFines(c *gin.Context)
	GetUnpaidFines(c *gin.Context)
	GetAllFines(c *gin.Context)
}

type Handler struct {
	srv service.AuthService
	ii  service.IIService
	log *slog.Logger
}

func NewHandler(log *slog.Logger, sr service.AuthService, II service.IIService) Handlers {
	return &Handler{log: log, srv: sr, ii: II}
}

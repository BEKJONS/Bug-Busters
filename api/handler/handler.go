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

	GetProfile(c *gin.Context)
	AddImage(c *gin.Context)
	GetImage(c *gin.Context)
	GetPaidFinesU(c *gin.Context)
	GetUnpaid(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type Handler struct {
	auth service.AuthService
	ii   service.IIService
	user service.UserService

	log *slog.Logger
}

func NewHandler(log *slog.Logger, sr service.AuthService, II service.IIService, u service.UserService) Handlers {
	return &Handler{log: log, auth: sr, ii: II, user: u}
}

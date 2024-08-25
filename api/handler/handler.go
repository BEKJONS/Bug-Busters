package handler

import (
	"bug_busters/internal/service"
	"bug_busters/pkg/models"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type IIHandler interface {
	CreateFines(fine *models.FineReq) error
	AcceptFinesById(accept models.FineAccept) error
	GetPaidFines(pagination models.Pagination) (*models.Fines, error)
	GetUnpaidFines(pagination models.Pagination) (*models.Fines, error)
	GetAllFines(pagination models.Pagination) (*models.Fines, error)
}

type Handler struct {
	srv service.AuthService
	ii  service.IIService
	log *slog.Logger
}

func NewAuthHandler(log *slog.Logger, sr service.AuthService, II service.IIService) AuthHandler {
	return &Handler{log: log, srv: sr, ii: II}
}

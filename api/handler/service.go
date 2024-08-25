package handler

import (
	"net/http"
	models "bug_busters/pkg/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all services
// @Description Get all services
// @Tags service
// @Accept json
// @Produce json
// @Success 200 {object} models.Services
// @Failure 400 {object} string
// @Router /service [get]
func (h *Handler) GetAllServices(c *gin.Context) {
	services, err := h.serv.GetServices()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}


// @Summary Get service
// @Description Get service
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Service
// @Failure 400 {object} string
// @Router /service/{id} [get]
func (h *Handler) GetService(c *gin.Context) {
	id := c.Param("id")
	service, err := h.serv.GetService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

// @Summary Create service
// @Description Create service
// @Tags service
// @Accept json
// @Produce json
// @Param service body models.Service true "service"
// @Success 200 {object} models.Service
// @Failure 400 {object} string
// @Router /service [post]
func (h *Handler) CreateService(c *gin.Context) {
	service := &models.Service{}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := h.serv.CreateService(service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

// @Summary Update service
// @Description Update service
// @Tags service
// @Accept json
// @Produce jsonGetAllServices
// @Param id path string true "id"
// @Param service body models.Service true "service"
// @Success 200 {object} models.Service
// @Failure 400 {object} string
// @Router /service/update [put]
func (h *Handler) UpdateService(c *gin.Context) {
	service := &models.Service{}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := h.serv.UpdateService(service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

// @Summary Delete service
// @Description Delete service
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Message
// @Failure 400 {object} string
// @Router /service/{id} [delete]
func (h *Handler) DeleteService(c *gin.Context) {
	id := c.Param("id")
	service, err := h.serv.DeleteService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}
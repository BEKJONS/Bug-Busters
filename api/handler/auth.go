package handler

import (
	"bug_busters/pkg/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "bug_busters/api/docs"
)

// Register godoc
// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var auth *models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	err := h.auth.Register(*auth)
	if err != nil {
		h.log.Error("Error occurred while register", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: "Successfully registered"})
}

// Login godoc
// @Summary LoginEmail Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var auth *models.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	res, err := h.auth.Login(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AddLicense godoc
// @Summary Add a new license
// @Description Add a new license to the system
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param License body models.LicenceNumber true "License information"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /auth/add_license [post]
func (h *Handler) AddLicense(c *gin.Context) {
	var req models.LicenceNumber

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		log.Println(err)
		return
	}

	log.Println(req)

	res, err := h.auth.AddLicence(&req)
	if err != nil {
		h.log.Error("Error occurred while adding license", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

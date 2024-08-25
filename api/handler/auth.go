package handler

import (
	"bug_busters/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /register [post]
func (h *authHandler) Register(c *gin.Context) {
	var auth *models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.Register(*auth)
	if err != nil {
		h.log.Error("Error occurred while register", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully registered"})
}

// Login godoc
// @Summary LoginEmail Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.LoginEmailResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login/email [post]
func (h *authHandler) Login(c *gin.Context) {
	var auth *models.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.Login(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

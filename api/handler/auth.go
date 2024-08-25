package handler

import (
	"bug_busters/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "bug_busters/api/docs"
)

// Register godoc
// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /register [post]
func (h *authHandler) Register(c *gin.Context) {
	var auth *models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	err := h.srv.Register(*auth)
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
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login [post]
func (h *authHandler) Login(c *gin.Context) {
	var auth *models.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	res, err := h.srv.Login(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

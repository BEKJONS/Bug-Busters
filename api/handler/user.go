package handler

import (
	"bug_busters/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.UserProfile
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	id := c.MustGet("user_id").(string)
	profile, err := h.user.GetProfile(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get user profile", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetProfileAdmin godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.UserProfile
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/profile/{id} [get]
func (h *Handler) GetProfileAdmin(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.user.GetProfile(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get user profile", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// AddImage godoc
// @Summary Add Car Image
// @Description Upload a new car image
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param image body models.UpdateCarImage true "Car Image"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/image [post]
func (h *Handler) AddImage(c *gin.Context) {
	var req models.UpdateCarImage
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	msg, err := h.user.AddImage(&req)
	if err != nil {
		h.log.Error("Failed to add image", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

// GetImage godoc
// @Summary Get Car Image
// @Description Retrieve the image of a user's car
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Url
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/image [get]
func (h *Handler) GetImage(c *gin.Context) {
	id := c.Param("id")
	image, err := h.user.GetImage(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get image", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, image)
}

// GetPaidFinesU godoc
// @Summary Get Paid Fines
// @Description Retrieve all paid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/paid_fines [get]
func (h *Handler) GetPaidFinesU(c *gin.Context) {
	id := c.MustGet("id").(string)
	fines, err := h.user.GetPaidFinesU(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get paid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetPaidFinesAdmin godoc
// @Summary Get Paid Fines
// @Description Retrieve all paid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/paid_fines/{id} [get]
func (h *Handler) GetPaidFinesAdmin(c *gin.Context) {
	id := c.Param("id")
	fines, err := h.user.GetPaidFinesU(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get paid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetUnpaid godoc
// @Summary Get Unpaid Fines
// @Description Retrieve all unpaid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/unpaid_fines [get]
func (h *Handler) GetUnpaid(c *gin.Context) {
	id := c.MustGet("id").(string)
	fines, err := h.user.GetUnpaid(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get unpaid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetUnpaidAdmin godoc
// @Summary Get Unpaid Fines
// @Description Retrieve all unpaid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/unpaid_fines/{id} [get]
func (h *Handler) GetUnpaidAdmin(c *gin.Context) {
	id := c.Param("id")
	fines, err := h.user.GetUnpaid(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get unpaid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	msg, err := h.user.DeleteUser(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to delete user", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

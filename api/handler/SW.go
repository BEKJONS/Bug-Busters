package handler

import (
	"bug_busters/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLicense godoc
// @Summary Create License
// @Description Create a new driver license
// @Tags License
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param License body models.DriverLicense true "Create license"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /single_window/license [post]
func (h *Handler) CreateLicense(c *gin.Context) {
	var license models.DriverLicense

	if err := c.ShouldBindJSON(&license); err != nil {
		h.log.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	msg, err := h.sw.CreateLicense(license)
	if err != nil {
		h.log.Error("Error occurred while creating license", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}

// CreatePassport godoc
// @Summary Create Passport
// @Description Create a new passport
// @Tags Passport
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Passport body models.CardId true "Create passport"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /single_window/passport [post]
func (h *Handler) CreatePassport(c *gin.Context) {
	var passport models.CardId

	if err := c.ShouldBindJSON(&passport); err != nil {
		h.log.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	msg, err := h.sw.CreatePassport(passport)
	if err != nil {
		h.log.Error("Error occurred while creating passport", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}

// GetAllLicenses godoc
// @Summary Get All Licenses
// @Description Retrieve all driver licenses
// @Tags License
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []models.DriverLicense
// @Failure 500 {object} models.Error
// @Router /single_window/licenses [get]
func (h *Handler) GetAllLicenses(c *gin.Context) {
	licenses, err := h.sw.GetLicenseAll()
	if err != nil {
		h.log.Error("Error occurred while retrieving licenses", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, licenses)
}

// GetAllPassports godoc
// @Summary Get All Passports
// @Description Retrieve all passports
// @Tags Passport
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.PassportsId
// @Failure 500 {object} models.Error
// @Router /single_window/passports [get]
func (h *Handler) GetAllPassports(c *gin.Context) {
	passports, err := h.sw.GetPassportAll()
	if err != nil {
		h.log.Error("Error occurred while retrieving passports", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, passports)
}

// DeleteLicense godoc
// @Summary Delete License
// @Description Delete a driver license by license number
// @Tags License
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param LicenseNumber body models.LicenceNumbers true "Delete license"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /single_window/license [delete]
func (h *Handler) DeleteLicense(c *gin.Context) {
	var license models.LicenceNumbers

	if err := c.ShouldBindJSON(&license); err != nil {
		h.log.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	msg, err := h.sw.DeleteLicense(license)
	if err != nil {
		h.log.Error("Error occurred while deleting license", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}

// DeletePassport godoc
// @Summary Delete Passport
// @Description Delete a passport by passport ID
// @Tags Passport
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param PassportId body models.PassportId true "Delete passport"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /single_window/passport [delete]
func (h *Handler) DeletePassport(c *gin.Context) {
	var passport models.PassportId

	if err := c.ShouldBindJSON(&passport); err != nil {
		h.log.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	msg, err := h.sw.DeletePassport(passport)
	if err != nil {
		h.log.Error("Error occurred while deleting passport", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}

package handler

import (
	"bug_busters/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFines godoc
// @Summary Create Fines
// @Description Create a new fine
// @Tags Fines
// @Accept json
// @Produce json
// @Param Fine body models.FineReq true "Create fine"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /fines [post]
func (h *Handler) CreateFines(c *gin.Context) {
	var fine models.FineReq

	if err := c.ShouldBindJSON(&fine); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	err := h.ii.CreateFines(&fine)
	if err != nil {
		h.log.Error("Error occurred while creating fine", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: "Fine created successfully"})
}

// AcceptFinesById godoc
// @Summary Accept Fines By ID
// @Description Accept a fine by updating its payment date
// @Tags Fines
// @Accept json
// @Produce json
// @Param FineAccept body models.FineAccept true "Accept fine"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /fines/:id/accept [put]
func (h *Handler) AcceptFinesById(c *gin.Context) {
	var accept models.FineAccept

	if err := c.ShouldBindJSON(&accept); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	err := h.ii.AcceptFinesById(accept)
	if err != nil {
		h.log.Error("Error occurred while accepting fine", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: "Fine accepted successfully"})
}

// GetPaidFines godoc
// @Summary Get Paid Fines
// @Description Retrieve all paid fines
// @Tags Fines
// @Accept json
// @Produce json
// @Param page query int true "Pagination"
// @Param limit query int true "Limit"
// @Success 200 {object} models.Fines
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /fines/paid [get]
func (h *Handler) GetPaidFines(c *gin.Context) {
	var pagination models.Pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagination.Page = page

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	pagination.Limit = limit

	fines, err := h.ii.GetPaidFines(pagination)
	if err != nil {
		h.log.Error("Error occurred while retrieving paid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fines)
}

// GetUnpaidFines godoc
// @Summary Get Unpaid Fines
// @Description Retrieve all unpaid fines
// @Tags Fines
// @Accept json
// @Produce json
// @Param page query int true "Pagination"
// @Param limit query int true "Limit"
// @Success 200 {object} models.Fines
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /fines/unpaid [get]
func (h *Handler) GetUnpaidFines(c *gin.Context) {
	var pagination models.Pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagination.Page = page

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	pagination.Limit = limit

	fines, err := h.ii.GetUnpaidFines(pagination)
	if err != nil {
		h.log.Error("Error occurred while retrieving unpaid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fines)
}

// GetAllFines godoc
// @Summary Get All Fines
// @Description Retrieve all fines
// @Tags Fines
// @Accept json
// @Produce json
// @Param Pagination query models.Pagination true "Pagination"
// @Success 200 {object} models.Fines
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /fines [get]
func (h *Handler) GetAllFines(c *gin.Context) {
	var pagination models.Pagination

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagination.Page = page

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	pagination.Limit = limit

	fines, err := h.ii.GetAllFines(pagination)
	if err != nil {
		h.log.Error("Error occurred while retrieving all fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fines)
}

// SendAcceptation godoc
// @Summary Accept a fine by ID
// @Description Retrieve the ID of the accepted fine
// @Tags Fines
// @Accept json
// @Produce json
// @Success 200 {object} models.Message "Accepted fine ID"
// @Failure 400 {object} models.Error "Bad request"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /fines/send_acceptation [post]
func (h *Handler) SendAcceptation(c *gin.Context) {
	id := c.MustGet("id").(string)
	c.JSON(http.StatusOK, models.Message{Message: id})
}

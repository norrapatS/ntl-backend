package handlers

import (
	"net/http"

	"ntl-test/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(
	adminService *services.AdminService,
) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) GetApplications(c *gin.Context) {
	applications, err := h.adminService.GetApplications()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get applications",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": applications,
	})
}

func (h *AdminHandler) GetApplicationByTransactionNumber(c *gin.Context) {
	transactionNo := c.Param("transactionNo")

	if transactionNo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction number is required",
		})
		return
	}

	application, err := h.adminService.GetApplicationByTransactionNo(transactionNo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "application not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": application,
	})
}

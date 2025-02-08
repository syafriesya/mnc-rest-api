package user

import (
	"mnc-rest-api/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserController) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")

	var req domain.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "FAILED", "message": "Invalid request"})
		return
	}

	response, err := h.userUsecase.UpdateProfile(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILED", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": response,
	})
}

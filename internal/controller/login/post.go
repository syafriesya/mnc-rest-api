package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *LoginController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		Pin         string `json:"pin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := l.loginUsecase.LoginUser(ctx, req.PhoneNumber, req.Pin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "Fail",
			"message": "Invalid phone number or pin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"result": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

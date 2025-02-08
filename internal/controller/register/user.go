package register

import (
	"fmt"
	"mnc-rest-api/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *RegisterController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("req", req)
	resp, err := r.registerUsecase.RegisterUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Success",
		"result": resp,
	})
}

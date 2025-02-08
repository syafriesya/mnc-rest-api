package topup

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TopupController) TopUp(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetString("userID")

	fmt.Println("controller", userID)
	resp, err := t.topupUsecase.TopUp(ctx, userID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"result": resp,
	})
}

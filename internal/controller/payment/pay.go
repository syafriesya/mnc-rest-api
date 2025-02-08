package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p PaymentController) Pay(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Amount  float64 `json:"amount"`
		Remarks string  `json:"remarks"`
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

	resp, err := p.paymentUsecase.Pay(ctx, userID, req.Amount, req.Remarks)
	if err != nil {
		if err.Error() == "balance is not enough" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Fail",
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"result": resp,
	})
}

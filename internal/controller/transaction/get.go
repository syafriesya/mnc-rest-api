package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TransactionController) GetTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetString("userID")

	transactions, err := t.transactionUsecase.GetUserTransactions(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transactions,
	})
}

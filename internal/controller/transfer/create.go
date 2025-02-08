package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TransferController) Transfer(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		TargetUser string  `json:"target_user"`
		Amount     float64 `json:"amount"`
		Remarks    string  `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 || req.TargetUser == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	senderID := c.GetString("userID")

	resp, err := t.transferUsecase.Transfer(ctx, senderID, req.TargetUser, req.Amount, req.Remarks)
	if err != nil {
		if err.Error() == "Balance is not enough" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Fail",
				"message": "Balance is not enough",
			})
		} else if err.Error() == "Target user not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Fail",
				"message": "Target user not found",
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

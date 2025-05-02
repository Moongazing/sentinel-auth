package user

import (
	"net/http"

	"SentinelAuth/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}

	userID, err := utils.VerifyEmailToken(h.Service.Rdb, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired token"})
		return
	}

	if err := h.Service.MarkEmailVerified(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email verified successfully"})
}

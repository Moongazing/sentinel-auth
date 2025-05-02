package user

import (
	"net/http"

	"SentinelAuth/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendOTP(c *gin.Context) {
	userID := c.GetUint("userID")

	otp, err := utils.GenerateOTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate OTP"})
		return
	}

	if err := utils.StoreOTP(h.Service.Rdb, userID, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store OTP"})
		return
	}

	// Not real email, just a placeholder
	// In production, send this via SMTP
	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent (simulated)",
		"otp":     otp, // gösteriyoruz, çünkü e-posta sistemi henüz devrede değil
	})
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	userID := c.GetUint("userID")

	var body struct {
		OTP string `json:"otp" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.VerifyOTP(h.Service.Rdb, userID, body.OTP) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
}

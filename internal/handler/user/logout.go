package user

import (
	"SentinelAuth/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type LogoutHandler struct {
	Redis *redis.Client
}

func NewLogoutHandler(redis *redis.Client) *LogoutHandler {
	return &LogoutHandler{Redis: redis}
}

func (l *LogoutHandler) Logout(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Varsayılan süre 7 gün (refresh token süresi)
	blacklistDuration := 7 * 24 * time.Hour

	if err := utils.BlacklistRefreshToken(l.Redis, body.RefreshToken, blacklistDuration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

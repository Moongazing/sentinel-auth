package middleware

import (
	"SentinelAuth/internal/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleRequired(roleName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		var result struct {
			RoleID *uint
		}

		err := infrastructure.DB.
			Table("users").
			Select("role_id").
			Where("id = ?", userID).
			Scan(&result).Error

		if err != nil || result.RoleID == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		var role struct {
			Name string
		}

		err = infrastructure.DB.
			Table("roles").
			Select("name").
			Where("id = ?", *result.RoleID).
			Scan(&role).Error

		if err != nil || role.Name != roleName {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

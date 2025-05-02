package user

import (
	"SentinelAuth/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *usecase.UserService
}

func NewHandler(service *usecase.UserService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Register(req.Email, req.Password, req.FullName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

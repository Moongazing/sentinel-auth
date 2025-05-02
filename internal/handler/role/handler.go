package role

import (
	"net/http"
	"strconv"

	"SentinelAuth/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *usecase.RoleService
}

func NewHandler(service *usecase.RoleService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.Service.CreateRole(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, RoleResponse{ID: role.ID, Name: role.Name})
}

func (h *Handler) GetAll(c *gin.Context) {
	roles, _ := h.Service.GetAllRoles()
	var out []RoleResponse
	for _, r := range roles {
		out = append(out, RoleResponse{ID: r.ID, Name: r.Name})
	}
	c.JSON(http.StatusOK, out)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := h.Service.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}
	c.JSON(http.StatusOK, RoleResponse{ID: role.ID, Name: role.Name})
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.Service.UpdateRole(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, RoleResponse{ID: role.ID, Name: role.Name})
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
}

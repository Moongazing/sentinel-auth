package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAll(c *gin.Context) {
	var query PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil || query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}

	users, total, err := h.Service.GetUsers(query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var items []UserListItemDTO
	for _, u := range users {
		items = append(items, UserListItemDTO{
			ID:       u.ID,
			Email:    u.Email,
			FullName: u.FullName,
		})
	}

	c.JSON(http.StatusOK, PaginatedUsersResponse{
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalCount: total,
		Data:       items,
	})
}

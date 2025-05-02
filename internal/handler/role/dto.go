package role

type CreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

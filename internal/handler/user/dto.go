package user

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type PaginationQuery struct {
	Page     int `form:"page" binding:"gte=1"`
	PageSize int `form:"page_size" binding:"gte=1,lte=100"`
}

type PaginatedUsersResponse struct {
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalCount int64             `json:"total_count"`
	Data       []UserListItemDTO `json:"data"`
}

type UserListItemDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
type AssignRoleRequest struct {
	RoleID uint `json:"role_id" binding:"required"`
}

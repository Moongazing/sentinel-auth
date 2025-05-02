package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	rolehandler "SentinelAuth/internal/handler/role"
	"SentinelAuth/internal/handler/user"

	"SentinelAuth/internal/middleware"
	"SentinelAuth/internal/repository"
	"SentinelAuth/internal/usecase"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// ==== USER DEPENDENCIES ====
	userRepo := repository.NewUserRepository(db)
	userService := &usecase.UserService{
		Repo: userRepo,
		Rdb:  rdb,
	}
	userHandler := user.NewHandler(userService)
	logoutHandler := user.NewLogoutHandler(rdb)

	// ==== ROLE DEPENDENCIES ====
	roleRepo := repository.NewRoleRepository(db)
	roleService := &usecase.RoleService{Repo: roleRepo}
	roleHandler := rolehandler.NewHandler(roleService)

	// ==== PUBLIC ROUTES ====
	api := router.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.POST("/refresh", userHandler.RefreshToken)
	}

	// ==== PROTECTED ROUTES (JWT Middleware) ====
	authGroup := router.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/users", userHandler.GetAll)
		authGroup.POST("/logout", logoutHandler.Logout)

		// ROLE CRUD
		authGroup.POST("/roles", roleHandler.Create)
		authGroup.GET("/roles", roleHandler.GetAll)
		authGroup.GET("/roles/:id", roleHandler.GetByID)
		authGroup.PUT("/roles/:id", roleHandler.Update)
		authGroup.DELETE("/roles/:id", roleHandler.Delete)
		authGroup.PUT("/users/:id/role", userHandler.AssignRoleToUser)
		authGroup.POST("/otp/send", userHandler.SendOTP)
		authGroup.POST("/otp/verify", userHandler.VerifyOTP)
		api.GET("/verify-email", userHandler.VerifyEmail)

	}
}

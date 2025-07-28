package router

import (
	"gin-template/controller"
	"gin-template/docs"
	"gin-template/middleware"
	"gin-template/models"
	"gin-template/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	// Create Gin engine
	r := gin.Default()

	// Global middleware
	r.Use(middleware.CORS())

	// Initialize services
	userService := service.NewUserService(db)

	// Initialize controllers
	userController := controller.NewUserController(userService)

	// API route group
	api := r.Group("/api/v1")

	// User routes - using new middleware architecture
	userRoutes := api.Group("/users")
	{
		// POST /users - create user, auto-bind JSON body
		userRoutes.POST("",
			middleware.BindAndCall(
				userController.CreateUser,
				(*models.CreateUserRequest)(nil),
			))

		// GET /users - get user list, auto-bind query parameters
		userRoutes.GET("",
			middleware.BindAndCall(
				userController.GetUsers,
				(*models.GetUsersQuery)(nil),
			))

		// GET /users/:id - get single user
		userRoutes.GET("/:id", userController.GetUser)

		// PUT /users/:id - update user, auto-bind JSON body
		userRoutes.PUT("/:id",
			middleware.BindAndCall(
				userController.UpdateUser,
				(*models.UpdateUserRequest)(nil),
			))

		// DELETE /users/:id - delete user
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		middleware.SuccessResponse(c, gin.H{
			"status":  "healthy",
			"message": "API is running",
		})
	})

	return r
}

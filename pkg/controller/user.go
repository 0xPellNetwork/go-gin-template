package controller

import (
	"net/http"
	"strconv"

	"gin-template/pkg/middleware"
	"gin-template/pkg/models"
	"gin-template/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User creation data"
// @Success 200 {object} middleware.Response{data=models.User} "User created successfully"
// @Failure 400 {object} middleware.Response "Invalid request body"
// @Failure 500 {object} middleware.Response "Internal server error"
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context, req models.CreateUserRequest) {
	user, err := uc.userService.CreateUser(&req)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.SuccessResponse(c, user)
}

// GetUser retrieves a single user by ID
// @Summary Get user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.Response{data=models.User} "User found"
// @Failure 400 {object} middleware.Response "Invalid user ID"
// @Failure 404 {object} middleware.Response "User not found"
// @Router /users/{id} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	idStr, err := middleware.GetPathID(c)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := uc.userService.GetUser(uint(id))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	middleware.SuccessResponse(c, user)
}

// GetUsers retrieves a list of users with pagination
// @Summary Get users list
// @Description Get a paginated list of users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param name query string false "Filter by name"
// @Param email query string false "Filter by email"
// @Success 200 {object} middleware.Response{data=object} "Users retrieved successfully"
// @Failure 400 {object} middleware.Response "Invalid query parameters"
// @Failure 500 {object} middleware.Response "Internal server error"
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context, query models.GetUsersQuery) {
	users, total, err := uc.userService.GetUsers(&query)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]any{
		"users":     users,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	}

	middleware.SuccessResponse(c, response)
}

// UpdateUser updates an existing user
// @Summary Update user
// @Description Update an existing user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User update data"
// @Success 200 {object} middleware.Response{data=models.User} "User updated successfully"
// @Failure 400 {object} middleware.Response "Invalid request"
// @Failure 500 {object} middleware.Response "Internal server error"
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context, req models.UpdateUserRequest) {
	idStr, err := middleware.GetPathID(c)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := uc.userService.UpdateUser(uint(id), &req)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.SuccessResponse(c, user)
}

// DeleteUser deletes a user by ID
// @Summary Delete user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.Response "User deleted successfully"
// @Failure 400 {object} middleware.Response "Invalid user ID"
// @Failure 500 {object} middleware.Response "Internal server error"
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	idStr, err := middleware.GetPathID(c)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err = uc.userService.DeleteUser(uint(id))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	middleware.SuccessResponse(c, nil)
}

package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

// List godoc
// @Summary List users
// @Description Get a list of active users
// @Tags users
// @Security Bearer
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users [get]
func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	// TODO: Fetch users from database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"users": []gin.H{
			{
				"id":       1,
				"email":    "user1@example.com",
				"username": "user1",
			},
			{
				"id":       2,
				"email":    "user2@example.com",
				"username": "user2",
			},
		},
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       2,
			"total_pages": 1,
			"offset":      offset,
		},
	})
}

// Get godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func (h *UserHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// TODO: Fetch user from database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":         id,
			"email":      "user@example.com",
			"username":   "user",
			"full_name":  "John Doe",
			"is_active":  true,
			"created_at": "2024-01-01T00:00:00Z",
			"updated_at": "2024-01-01T00:00:00Z",
		},
	})
}

// Update godoc
// @Summary Update user
// @Description Update user details
// @Tags users
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body map[string]interface{} true "User update details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update user in database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data": gin.H{
			"id": id,
		},
	})
}

// Delete godoc
// @Summary Delete user
// @Description Delete a user account
// @Tags users
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// TODO: Delete user from database using sqlc

	c.Status(http.StatusNoContent)
	_ = id
}
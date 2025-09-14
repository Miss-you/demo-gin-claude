package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	db *sql.DB
}

func NewPostHandler(db *sql.DB) *PostHandler {
	return &PostHandler{db: db}
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status"`
}

type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
}

// List godoc
// @Summary List posts
// @Description Get a list of published posts
// @Tags posts
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /posts [get]
func (h *PostHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	// TODO: Fetch posts from database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"posts": []gin.H{
			{
				"id":           1,
				"user_id":      1,
				"title":        "First Post",
				"content":      "This is the first post",
				"status":       "published",
				"published_at": "2024-01-01T00:00:00Z",
			},
			{
				"id":           2,
				"user_id":      2,
				"title":        "Second Post",
				"content":      "This is the second post",
				"status":       "published",
				"published_at": "2024-01-02T00:00:00Z",
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
// @Summary Get post by ID
// @Description Get post details by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /posts/{id} [get]
func (h *PostHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// TODO: Fetch post from database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           id,
			"user_id":      1,
			"title":        "Sample Post",
			"content":      "This is a sample post content",
			"status":       "published",
			"published_at": "2024-01-01T00:00:00Z",
			"created_at":   "2024-01-01T00:00:00Z",
			"updated_at":   "2024-01-01T00:00:00Z",
		},
	})
}

// Create godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Security Bearer
// @Accept json
// @Produce json
// @Param request body CreatePostRequest true "Post details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /posts [post]
func (h *PostHandler) Create(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if req.Status == "" {
		req.Status = "draft"
	}

	// TODO: Create post in database using sqlc

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"data": gin.H{
			"id":      1,
			"user_id": userID,
			"title":   req.Title,
			"content": req.Content,
			"status":  req.Status,
		},
	})
}

// Update godoc
// @Summary Update post
// @Description Update post details
// @Tags posts
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param request body UpdatePostRequest true "Post update details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /posts/{id} [put]
func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update post in database using sqlc

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"data": gin.H{
			"id": id,
		},
	})
}

// Delete godoc
// @Summary Delete post
// @Description Delete a post
// @Tags posts
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 204
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /posts/{id} [delete]
func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// TODO: Delete post from database using sqlc

	c.Status(http.StatusNoContent)
	_ = id
}
package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreatePost creates a new post
// @Summary Create Post
// @Description Create a new post
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param post body pb.PostCreate true "Post data"
// @Success 200 {string} string "Post created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/create [post]
func (h *Handler) CreatePost(c *gin.Context) {
	var req pb.PostCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Post.CreatePost(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create post:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Post created successfully")
}

// GetPost retrieves a post by ID
// @Summary Get Post
// @Description Retrieve a post by its ID
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Success 200 {object} pb.PostGet "Post details"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Post not found"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/{id} [get]
func (h *Handler) GetPost(c *gin.Context) {
	postID := c.Param("id")

	req := &pb.ById{Id: postID}
	res, err := h.Clients.Post.GetPost(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get post:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdatePost updates an existing post
// @Summary Update Post
// @Description Update an existing post
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Param post body pb.PostUpdate true "Post data"
// @Success 200 {string} string "Post updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/{id} [put]
func (h *Handler) UpdatePost(c *gin.Context) {
	postID := c.Param("id")

	var body pb.PostUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	req := pb.PostUpdate{
		Id:   postID,
		Body: &body,
	}

	_, err := h.Clients.Post.UpdatePost(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update post:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Post updated successfully")
}

// DeletePost deletes a post by ID
// @Summary Delete Post
// @Description Delete a post by its ID
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Post ID"
// @Success 200 {string} string "Post deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Post not found"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/{id} [delete]
func (h *Handler) DeletePost(c *gin.Context) {
	postID := c.Param("id")

	req := &pb.ById{Id: postID}
	_, err := h.Clients.Post.DeletePost(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete post:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Post deleted successfully")
}

// GetPosts retrieves a list of posts with pagination
// @Summary Get Posts
// @Description Retrieve a list of posts with pagination
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.PostList "List of posts"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/list [get]
func (h *Handler) GetPosts(c *gin.Context) {
	var filter pb.Pagination

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = int64(l)
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Offset = int64(o)
		}
	}

	res, err := h.Clients.Post.GetPosts(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list books:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}




// AddPostPicture adds a picture to a post
// @Summary Add Post Picture
// @Description Add a picture to a post
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param picture body pb.PostPicture true "Post picture data"
// @Success 200 {string} string "Picture added successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/picture/add [post]
func (h *Handler) AddPicturePost(c *gin.Context) {
	var req pb.PostPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Post.AddPostPicture(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to add post picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture added successfully")
}

// DeletePostPicture deletes a picture from a post
// @Summary Delete Post Picture
// @Description Delete a picture from a post
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param picture body pb.PostPicture true "Post picture data"
// @Success 200 {string} string "Picture deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /posts/picture/delete [post]
func (h *Handler) DeletePicturePost(c *gin.Context) {
	var req pb.PostPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Post.DeletePostPicture(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete post picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture deleted successfully")
}

package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateCategory creates a new category
// @Summary 		Create Category
// @Description 	Create a new category
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		category body    pb.CategoryCreateReq  true   "Category data"
// @Success 		200  {string}  string "Category created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/create [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var req pb.CategoryCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	_, err := h.Clients.Category.CreateCategory(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Category created successfully")
}

// UpdateCategory updates an existing category
// @Summary 		Update Category
// @Description 	Update an existing category
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Param     		category    body    pb.CategoryUpt  true   "Category data"
// @Success 		200  {string}  string "Category updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var body pb.CategoryUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}
	req := pb.CategoryUpdateReq{
		Id:   categoryID,
		Body: &body,
	}

	_, err := h.Clients.Category.UpdateCategory(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Category updated successfully")
}

// GetCategory retrieves a category by ID
// @Summary 		Get Category
// @Description 	Retrieve a category by its ID
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Success 		200  {object}  pb.Category "Category details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Category not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	categoryID := c.Param("id")

	req := &pb.ById{Id: categoryID}

	res, err := h.Clients.Category.GetCategory(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListCategories retrieves a list of categories
// @Summary List Categories
// @Description Retrieve a list of categories
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.CategoryListRes "List of categories"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/list [get]
func (h *Handler) ListCategories(c *gin.Context) {
	// Initialize filter and set default values
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

	// If not in cache, retrieve from gRPC
	res, err := h.Clients.Category.ListCategories(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list categories:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// DeleteCategory deletes a category by ID
// @Summary 		Delete Category
// @Description 	Delete a category by its ID
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Success 		200  {string}  string "Category deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Category not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	req := &pb.ById{Id: categoryID}

	_, err := h.Clients.Category.DeleteCategory(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Category deleted successfully")
}

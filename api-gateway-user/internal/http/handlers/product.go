package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateProduct creates a new product
// @Summary 		Create Product
// @Description 	Create a new product
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		product body    pb.ProductCreate  true   "Product data"
// @Success 		200  {string}  string "Product created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/create [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req pb.ProductCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Product.CreateProduct(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create product:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Product created successfully")
}

// UpdateProduct updates an existing product
// @Summary 		Update Product
// @Description 	Update an existing product
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Product ID"
// @Param     		product    body    pb.ProductUpt  true   "Product data"
// @Success 		200  {string}  string "Product updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var body pb.ProductUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}
	req := pb.ProductUpdate{
		Id:   productID,
		Body: &body,
	}

	_, err := h.Clients.Product.UpdateProduct(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update product:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Product updated successfully")
}

// DeleteProduct deletes a product by ID
// @Summary 		Delete Product
// @Description 	Delete a product by its ID
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Product ID"
// @Success 		200  {string}  string "Product deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Product not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	req := &pb.ById{Id: productID}

	_, err := h.Clients.Product.DeleteProduct(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete product:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Product deleted successfully")
}

// GetProduct retrieves a product by ID
// @Summary 		Get Product
// @Description 	Retrieve a product by its ID
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Product ID"
// @Success 		200  {object}  pb.ProductGet "Product details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Product not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	productID := c.Param("id")

	req := &pb.ById{Id: productID}

	res, err := h.Clients.Product.GetProduct(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get product:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListProducts retrieves a list of products
// @Summary List Products
// @Description Retrieve a list of products
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param title query string false "Title"
// @Param price_from query float32 false "Price from"
// @Param price_to query float32 false "Price to"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ProductList "List of products"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /products/list [get]
func (h *Handler) ListProducts(c *gin.Context) {
	var req pb.ProductFilter
	req.Pagination = &pb.Pagination{}

	req.Title = c.Query("title")

	if priceFrom := c.Query("price_from"); priceFrom != "" {
		if pf, err := strconv.ParseFloat(priceFrom, 32); err == nil {
			req.PriceFrom = float32(pf)
		}
	}
	if priceTo := c.Query("price_to"); priceTo != "" {
		if pt, err := strconv.ParseFloat(priceTo, 32); err == nil {
			req.PriceTo = float32(pt)
		}
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			req.Pagination.Limit = int64(l)
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			req.Pagination.Offset = int64(o)
		}
	}

	// Retrieve products from gRPC service
	res, err := h.Clients.Product.ListProducts(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list products:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// AddPicture adds a picture to a product
// @Summary 		Add Picture
// @Description 	Add a picture to a product
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		picture body    pb.ProductPicture  true   "Picture data"
// @Success 		200  {string}  string "Picture added successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/pictures/add [post]
func (h *Handler) AddPictureProduct(c *gin.Context) {
	var req pb.ProductPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Product.AddPicture(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to add picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture added successfully")
}

// DeletePicture deletes a picture from a product
// @Summary 		Delete Picture
// @Description 	Delete a picture from a product
// @Tags 			Product
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		picture body    pb.ProductPicture  true   "Picture data"
// @Success 		200  {string}  string "Picture deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/products/pictures/delete [post]
func (h *Handler) DeletePictureProduct(c *gin.Context) {
	var req pb.ProductPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Product.DeletePicture(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture deleted successfully")
}

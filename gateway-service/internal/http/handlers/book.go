package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateBook creates a new book
// @Summary      Create Book
// @Description  Create a new book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        book body     pb.BookCreate  true  "Book data"
// @Success      200  {string} string         "Book created successfully"
// @Failure      400  {string} string         "Invalid request"
// @Failure      500  {string} string         "Internal server error"
// @Router       /books/create [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var req pb.BookCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Book.CreateBook(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create book:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Book created successfully")
}

// UpdateBook updates an existing book
// @Summary      Update Book
// @Description  Update an existing book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path     string       true   "Book ID"
// @Param        book body     pb.BookUpt   true   "Book data"
// @Success      200  {string} string       "Book updated successfully"
// @Failure      400  {string} string       "Invalid request"
// @Failure      500  {string} string       "Internal server error"
// @Router       /books/update/{id} [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	bookID := c.Param("id")

	var body pb.BookUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	req := pb.BookUpdate{
		Id:   bookID,
		Body: &body,
	}

	_, err := h.Clients.Book.UpdateBook(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update book:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Book updated successfully")
}

// DeleteBook deletes a book by ID
// @Summary      Delete Book
// @Description  Delete a book by its ID
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path     string       true   "Book ID"
// @Success      200  {string} string       "Book deleted successfully"
// @Failure      400  {string} string       "Invalid request"
// @Failure      500  {string} string       "Internal server error"
// @Router       /books/delete/{id} [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	req := &pb.ById{Id: bookID}

	_, err := h.Clients.Book.DeleteBook(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete book:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Book deleted successfully")
}

// GetBook retrieves a book by ID
// @Summary      Get Book
// @Description  Retrieve a book by its ID
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path     string       true   "Book ID"
// @Success      200  {object} pb.BookGet   "Book details"
// @Failure      400  {string} string       "Invalid request"
// @Failure      500  {string} string       "Internal server error"
// @Router       /books/{id} [get]
func (h *Handler) GetBook(c *gin.Context) {
	bookID := c.Param("id")

	req := &pb.ById{Id: bookID}

	res, err := h.Clients.Book.GetBook(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get book:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListBooks retrieves a list of books
// @Summary      List Books
// @Description  Retrieve a list of books
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        title       query   string  false  "Title"
// @Param        price_from  query   float64 false  "Price From"
// @Param        price_to    query   float64 false  "Price To"
// @Param        limit       query   int     false  "Limit"
// @Param        offset      query   int     false  "Offset"
// @Success      200  {object} pb.BookList  "List of books"
// @Failure      400  {string} string        "Invalid request"
// @Failure      500  {string} string        "Internal server error"
// @Router       /books/list [get]
func (h *Handler) ListBooks(c *gin.Context) {
	var filter pb.BookFilter

	if title := c.Query("title"); title != "" {
		filter.Title = title
	}

	if priceFrom := c.Query("price_from"); priceFrom != "" {
		if p, err := strconv.ParseFloat(priceFrom, 64); err == nil {
			filter.PriceFrom = float32(p)  // Convert float64 to float32
		}
	}

	if priceTo := c.Query("price_to"); priceTo != "" {
		if p, err := strconv.ParseFloat(priceTo, 64); err == nil {
			filter.PriceTo = float32(p)  // Convert float64 to float32
		}
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Pagination.Limit = int64(l)
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Pagination.Offset = int64(o)
		}
	}

	res, err := h.Clients.Book.ListBooks(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list books:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}


// AddPicture adds a picture to a book
// @Summary      Add Picture
// @Description  Add a picture to a book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        bookPicture body     pb.BookPicture  true  "Book picture data"
// @Success      200  {string} string  "Picture added successfully"
// @Failure      400  {string} string  "Invalid request"
// @Failure      500  {string} string  "Internal server error"
// @Router       /books/picture/add [post]
func (h *Handler) AddPicture(c *gin.Context) {
	var req pb.BookPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Book.AddPicture(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to add picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture added successfully")
}

// DeletePicture deletes a picture from a book
// @Summary      Delete Picture
// @Description  Delete a picture from a book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        bookPicture body     pb.BookPicture  true  "Book picture data"
// @Success      200  {string} string  "Picture deleted successfully"
// @Failure      400  {string} string  "Invalid request"
// @Failure      500  {string} string  "Internal server error"
// @Router       /books/picture/delete [delete]
func (h *Handler) DeletePicture(c *gin.Context) {
	var req pb.BookPicture
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Book.DeletePicture(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Picture deleted successfully")
}

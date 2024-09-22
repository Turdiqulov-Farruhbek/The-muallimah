package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateOrder creates a new order
// @Summary 		Create Order
// @Description 	Create a new order
// @Tags 			Order
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		order body    pb.OrderCreateReq  true   "Order data"
// @Success 		200  {string}  string "Order created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/orders/create [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req pb.OrderCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	_, err := h.Clients.Order.CreateOrder(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create order:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Order created successfully")
}

// UpdateOrder updates an existing order
// @Summary 		Update Order
// @Description 	Update an existing order
// @Tags 			Order
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Order ID"
// @Param     		order    body    pb.OrderUpt  true   "Order data"
// @Success 		200  {string}  string "Order updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/orders/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")

	var body pb.OrderUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}
	req := pb.OrderUpdateReq{
		Id:   orderID,
		Body: &body,
	}

	_, err := h.Clients.Order.UpdateOrder(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update order:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Order updated successfully")
}

// GetOrder retrieves an order by ID
// @Summary 		Get Order
// @Description 	Retrieve an order by its ID
// @Tags 			Order
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Order ID"
// @Success 		200  {object}  pb.OrderGetRes "Order details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Order not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/orders/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	orderID := c.Param("id")

	req := &pb.OrderGetReq{Id: orderID}

	res, err := h.Clients.Order.GetOrder(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get order:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListOrders retrieves a list of orders
// @Summary List Orders
// @Description Retrieve a list of orders
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.OrderListsRes "List of orders"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /orders/list [get]
func (h *Handler) ListOrders(c *gin.Context) {
	// Initialize the request and set default values
	var req pb.OrderListsReq
	req.Filter = &pb.Pagination{}

	req.UserId = c.Query("user_id")

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			req.Filter.Limit = int64(l)
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			req.Filter.Offset = int64(o)
		}
	}

	// Retrieve orders from gRPC service
	res, err := h.Clients.Order.ListOrders(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list orders:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// DeleteOrder deletes an order by ID
// @Summary 		Delete Order
// @Description 	Delete an order by its ID
// @Tags 			Order
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Order ID"
// @Success 		200  {string}  string "Order deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Order not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/orders/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")

	req := &pb.OrderDeleteReq{Id: orderID}

	_, err := h.Clients.Order.DeleteOrder(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete order:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Order deleted successfully")
}

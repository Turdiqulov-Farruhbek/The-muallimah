package handlers

import (
	"context"
	"encoding/json"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create a new notification with the provided details
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body pb.NotificationCreate true "Notification details"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notification/create [post]
type NotificationCreateDTO struct {
    Title   string `json:"title"`
    Body    string `json:"body"`
    Sender  string `json:"sender"`
    // Boshqa maydonlarni qo'shing
}

func (h *Handler) CreateNotification(c *gin.Context) {
	var req pb.NotificationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request body")
		return
	}

	_, err := h.Clients.Notification.CreateNotification(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create notification:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	// DTO yaratish
	dto := NotificationCreateDTO{
		Title:  req.RecieverId,
		Body:   req.Message,
		Sender: req.SenderId,
		// Boshqa maydonlarni qo'shing
	}

	message, err := json.Marshal(dto)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal DTO:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	err = h.Producer.ProduceMessages("notifications", message)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(201, "Notification created and sent to Kafka successfully")
}


// UpdateNotification godoc
// @Summary Update notification details by ID
// @Description Update the details of a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Param notification body pb.NotificationUpt true "Notification details"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notification/update/{id} [put]
func (h *Handler) UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	var body pb.NotificationUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request body")
		return
	}

	req := &pb.NotificationUpdate{
		NotificationId: id,
		Body:           &body,
	}

	_, err := h.Clients.Notification.UpdateNotification(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update notification:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	// Kafka orqali xabar yuborish
	message, err := json.Marshal(req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	err = h.Producer.ProduceMessages("notifications", message)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification updated and sent to Kafka successfully")
}


// DeleteNotification godoc
// @Summary Delete a notification by ID
// @Description Delete a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 200 {object} pb.Void
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notification/delete/{id} [delete]
func (h *Handler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	req := &pb.ById{Id: id}

	_, err := h.Clients.Notification.DeleteNotification(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete notification:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification deleted successfully")
}

// GetNotifications godoc
// @Summary List notifications with filters
// @Description Retrieve a list of notifications with optional filters
// @Tags notifications
// @Accept json
// @Produce json
// @Param reciever_id query string false "Receiver ID"
// @Param status query string false "Notification Status"
// @Param sender_id query string false "Sender ID"
// @Param limit query int32 false "Limit for pagination"
// @Param offset query int32 false "Offset for pagination"
// @Success 200 {object} pb.NotificationList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notifications [get]
func (h *Handler) ListNotifications(c *gin.Context) {
	var filter pb.NotifFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.Logger.ERROR.Println("Failed to bind query params:", err)
		c.JSON(400, "Invalid query parameters")
		return
	}

	resp, err := h.Clients.Notification.GetNotifications(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get notifications:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, resp)
}

// GetNotification godoc
// @Summary Get a notification by ID
// @Description Retrieve a specific notification by its ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 200 {object} pb.NotificationGet
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notification/{id} [get]
func (h *Handler) GetNotification(c *gin.Context) {
	id := c.Param("id")
	req := &pb.ById{Id: id}

	resp, err := h.Clients.Notification.GetNotification(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get notification:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, resp)
}

// NotifyAll godoc
// @Summary Send a notification to all users
// @Description Send a notification message to all users in the specified group
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body pb.NotificationMessage true "Notification message"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /notification/notify-all [post]
func (h *Handler) NotifyAll(c *gin.Context) {
	var req pb.NotificationMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request body")
		return
	}

	_, err := h.Clients.Notification.NotifyAll(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to notify all users:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Notification sent to all users successfully")
}

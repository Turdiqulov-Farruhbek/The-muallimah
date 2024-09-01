package handlers

// import (
// 	"context"

// 	pb "muallimah-gateway/internal/pkg/genproto"

// 	"github.com/gin-gonic/gin"
// )

// // CreateNotification godoc
// // @Summary Create a new notification
// // @Description Create a new notification with the provided details
// // @Tags notifications
// // @Accept json
// // @Produce json
// // @Param notification body pb.NotificationCreate true "Notification details"
// // @Success 201 {object} pb.Void
// // @Failure 400 {object} string "Bad Request"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /notification/create [post]
// func (h *Handler) CreateNotification(c *gin.Context) {
// 	var req pb.NotificationCreate
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		h.Logger.ERROR.Println("Failed to bind request:", err)
// 		c.JSON(400, "Invalid request body")
// 		return
// 	}

// 	_, err := h.Clients.Notification.CreateNotification(context.Background(), &req)
// 	if err != nil {
// 		h.Logger.ERROR.Println("Failed to create notification:", err)
// 		c.JSON(500, "Internal server error: "+err.Error())
// 		return
// 	}

// 	c.JSON(201, "Notification created successfully")
// }

// // DeleteNotification godoc
// // @Summary Delete a notification by ID
// // @Description Delete a specific notification by its ID
// // @Tags notifications
// // @Accept json
// // @Produce json
// // @Param id path string true "Notification ID"
// // @Success 200 {object} pb.Void
// // @Failure 404 {object} string "Not Found"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /notification/delete/{id} [delete]
// func (h *Handler) DeleteNotification(c *gin.Context) {
// 	id := c.Param("id")
// 	req := &pb.ById{Id: id}

// 	_, err := h.Clients.Notification.DeleteNotification(context.Background(), req)
// 	if err != nil {
// 		h.Logger.ERROR.Println("Failed to delete notification:", err)
// 		c.JSON(500, "Internal server error: "+err.Error())
// 		return
// 	}

// 	c.JSON(200, "Notification deleted successfully")
// }

// // UpdateNotification godoc
// // @Summary Update notification details by ID
// // @Description Update the details of a specific notification by its ID
// // @Tags notifications
// // @Accept json
// // @Produce json
// // @Param id path string true "Notification ID"
// // @Param notification body pb.NotificationUpdate true "Notification details"
// // @Success 200 {object} pb.Void
// // @Failure 400 {object} string "Bad Request"
// // @Failure 404 {object} string "Not Found"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /notification/update/{id} [put]
// func (h *Handler) UpdateNotification(c *gin.Context) {
// 	id := c.Param("id")
// 	var body pb.NotificationUpt
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		h.Logger.ERROR.Println("Failed to bind request:", err)
// 		c.JSON(400, "Invalid request body")
// 		return
// 	}

// 	req := &pb.NotificationUpdate{
// 		NotificationId: id,
// 		Body:           &body,
// 	}

// 	_, err := h.Clients.Notification.UpdateNotification(context.Background(), req)
// 	if err != nil {
// 		h.Logger.ERROR.Println("Failed to update notification:", err)
// 		c.JSON(500, "Internal server error: "+err.Error())
// 		return
// 	}

// 	c.JSON(200, "Notification updated successfully")
// }

// // GetNotifications godoc
// // @Summary List notifications with filters
// // @Description Retrieve a list of notifications with optional filters
// // @Tags notifications
// // @Accept json
// // @Produce json
// // @Param reciever_id query string false "Receiver ID

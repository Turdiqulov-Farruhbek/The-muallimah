package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateCertificate creates a new certificate
// @Summary      Create Certificate
// @Description  Create a new certificate
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path     string  true  "ID"
// @Success      200  {string} string "Certificate created successfully"
// @Failure      400  {string} string "Invalid request"
// @Failure      500  {string} string "Internal server error"
// @Router       /certificates/create [post]
func (h *Handler) CreateCertificate(c *gin.Context) {
	var req pb.ById
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Certificate.CreateCertificate(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create certificate:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Certificate created successfully")
}

// GetCertificate retrieves a certificate by ID
// @Summary      Get Certificate
// @Description  Retrieve a certificate by its ID
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path     string         true  "Certificate ID"
// @Success      200  {object} pb.CertificateGet "Certificate details"
// @Failure      400  {string} string         "Invalid request"
// @Failure      500  {string} string         "Internal server error"
// @Router       /certificates/{id} [get]
func (h *Handler) GetCertificate(c *gin.Context) {
	certificateID := c.Param("id")

	req := &pb.ById{Id: certificateID}

	res, err := h.Clients.Certificate.GetCertificate(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get certificate:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateCertificate updates an existing certificate
// @Summary      Update Certificate
// @Description  Update an existing certificate
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path     string             true  "Certificate ID"
// @Param        certificate body     pb.CertificateUpt true "Certificate data"
// @Success      200  {string} string           "Certificate updated successfully"
// @Failure      400  {string} string           "Invalid request"
// @Failure      500  {string} string           "Internal server error"
// @Router       /certificates/update/{id} [put]
func (h *Handler) UpdateCertificate(c *gin.Context) {
	certificateID := c.Param("id")

	var body pb.CertificateUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	req := pb.CertificateUpdate{
		Id:   certificateID,
		Body: &body,
	}

	// _, err := h.Clients.Certificate.UpdateCertificate(context.Background(), &req)
	// if err != nil {
	// 	h.Logger.ERROR.Println("Failed to update certificate:", err)
	// 	c.JSON(500, "Internal server error: "+err.Error())
	// 	return
	// }
	input,err := protojson.Marshal(&req)
	if err!= nil {
		h.Logger.ERROR.Println("Failed to marshal req:", err)
		c.JSON(500, "Invalid request: "+err.Error())
		return
    }
	if err := h.Producer.ProduceMessages("certificate-update",input); err != nil {
		h.Logger.ERROR.Println("Failed to produce message:", err)
        c.JSON(500, "Internal server error: "+err.Error())
        return
	}

	c.JSON(200, "Certificate updated successfully")
}

// DeleteCertificate deletes a certificate by ID
// @Summary      Delete Certificate
// @Description  Delete a certificate by its ID
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path     string  true  "Certificate ID"
// @Success      200  {string} string "Certificate deleted successfully"
// @Failure      400  {string} string "Invalid request"
// @Failure      500  {string} string "Internal server error"
// @Router       /certificates/delete/{id} [delete]
func (h *Handler) DeleteCertificate(c *gin.Context) {
	certificateID := c.Param("id")

	req := &pb.ById{Id: certificateID}

	_, err := h.Clients.Certificate.DeleteCertificate(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete certificate:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Certificate deleted successfully")
}

// ListCertificates retrieves a list of certificates
// @Summary      List Certificates
// @Description  Retrieve a list of certificates
// @Tags         Certificates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id   query    string  false  "User ID"
// @Param        course_id query    string  false  "Course ID"
// @Param        status    query    string  false  "Status"
// @Param        limit     query    int     false  "Limit"
// @Param        offset    query    int     false  "Offset"
// @Success      200  {object} pb.CertificateList "List of certificates"
// @Failure      400  {string} string              "Invalid request"
// @Failure      500  {string} string              "Internal server error"
// @Router       /certificates/list [get]
func (h *Handler) ListCertificates(c *gin.Context) {
	var filter pb.CertificateFilter
	filter.Filter = &pb.Pagination{}

	if userID := c.Query("user_id"); userID != "" {
		filter.UserId = userID
	}

	if courseID := c.Query("course_id"); courseID != "" {
		filter.CourseId = courseID
	}

	if status := c.Query("status"); status != "" {
		filter.Status = status
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Filter.Limit = int64(l)
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Filter.Offset = int64(o)
		}
	}

	res, err := h.Clients.Certificate.ListCertificates(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list certificates:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

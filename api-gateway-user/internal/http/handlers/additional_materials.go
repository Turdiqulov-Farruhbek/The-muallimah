package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateMaterial creates a new material
// @Summary 		Create Material
// @Description 	Create a new material
// @Tags 			Material
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		material body    pb.MaterialCreateReq  true   "Material data"
// @Success 		200  {string}  string "Material created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/materials/create [post]
func (h *Handler) CreateMaterial(c *gin.Context) {
	var req pb.MaterialCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	_, err := h.Clients.Material.CreateMaterial(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create material:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Material created successfully")
}

// UpdateMaterial updates an existing material
// @Summary 		Update Material
// @Description 	Update an existing material
// @Tags 			Material
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Material ID"
// @Param     		material    body    pb.MaterialUpdate  true   "Material data"
// @Success 		200  {string}  string "Material updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/materials/{id} [put]
func (h *Handler) UpdateMaterial(c *gin.Context) {
	materialID := c.Param("id")

	var body pb.MaterialUpdate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}
	req := pb.MaterialUpdateReq{
		Id:   materialID,
		Body: &body,
	}

	_, err := h.Clients.Material.UpdateMaterial(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update material:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Material updated successfully")
}

// GetMaterial retrieves a material by ID
// @Summary 		Get Material
// @Description 	Retrieve a material by its ID
// @Tags 			Material
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Material ID"
// @Success 		200  {object}  pb.MaterialGetRes "Material details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Material not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/materials/{id} [get]
func (h *Handler) GetMaterial(c *gin.Context) {
	materialID := c.Param("id")

	req := &pb.ById{Id: materialID}

	res, err := h.Clients.Material.GetMaterial(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get material:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListMaterials retrieves a list of materials
// @Summary List Materials
// @Description Retrieve a list of materials
// @Tags Material
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param lesson_id query string true "Lesson ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.MaterialListRes "List of materials"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /materials/list [get]
func (h *Handler) ListMaterials(c *gin.Context) {
	// Initialize the request and set default values
	var req pb.MaterialListReq
	req.Pagination = &pb.Pagination{}

	req.LessonId = c.Query("lesson_id")

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

	// Retrieve materials from gRPC service
	res, err := h.Clients.Material.ListMaterials(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list materials:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// DeleteMaterial deletes a material by ID
// @Summary 		Delete Material
// @Description 	Delete a material by its ID
// @Tags 			Material
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Material ID"
// @Success 		200  {string}  string "Material deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Material not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/materials/{id} [delete]
func (h *Handler) DeleteMaterial(c *gin.Context) {
	materialID := c.Param("id")

	req := &pb.ById{Id: materialID}

	_, err := h.Clients.Material.DeleteMaterial(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete material:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Material deleted successfully")
}

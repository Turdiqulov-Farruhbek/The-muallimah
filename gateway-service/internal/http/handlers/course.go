package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "muallimah-gateway/internal/pkg/genproto"
)


// CreateCourse godoc
// @Summary Create a new course
// @Description Create a new course with the provided details
// @Tags courses
// @Accept json
// @Produce json
// @Param course body pb.CourseCreateReq true "Course details"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /course/create [post]
func (h *Handler) CreateCourse(c *gin.Context) {
	var req pb.CourseCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request body")
		return
	}

	_, err := h.Clients.Course.CreateCourse(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(201, "Course created successfully")
}

// GetCourse godoc
// @Summary Get course details by ID
// @Description Retrieve details of a specific course by its ID
// @Tags courses
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} pb.Course
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /course/{id} [get]
func (h *Handler) GetCourse(c *gin.Context) {
	id := c.Param("id")
	req := &pb.ById{Id: id}

	res, err := h.Clients.Course.GetCourse(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateCourse godoc
// @Summary Update course details by ID
// @Description Update the details of a specific course by its ID
// @Tags courses
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Param course body pb.CourseUpt true "Course details"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /course/update/{id} [put]
func (h *Handler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var body pb.CourseUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request body")
		return
	}

	req := &pb.CourseUpdateReq{
		Id:   id,
		Body: &body,
	}

	_, err := h.Clients.Course.UpdateCourse(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Course updated successfully")
}

// DeleteCourse godoc
// @Summary Delete a course by ID
// @Description Delete a specific course by its ID
// @Tags courses
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} pb.Void
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /course/delete/{id} [delete]
func (h *Handler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	req := &pb.ById{Id: id}

	_, err := h.Clients.Course.DeleteCourse(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Course deleted successfully")
}

// ListCourses godoc
// @Summary List all courses
// @Description Retrieve a list of courses with optional filters
// @Tags courses
// @Accept json
// @Produce json
// @Param category_id query string false "Category ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.CourseListsRes
// @Failure 500 {object} string "Internal Server Error"
// @Router /course/list [get]
func (h *Handler) ListCourses(c *gin.Context) {
	var filter pb.CourseListsReq

	if categoryId := c.Query("category_id"); categoryId != "" {
		filter.CategoryId = categoryId
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

	res, err := h.Clients.Course.ListCourses(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list courses:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

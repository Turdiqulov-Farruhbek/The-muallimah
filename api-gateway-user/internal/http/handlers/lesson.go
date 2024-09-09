package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateLesson creates a new lesson
// @Summary 		Create Lesson
// @Description 	Create a new lesson
// @Tags 			Lesson
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		lesson body    pb.LessonCreateReq  true   "Lesson data"
// @Success 		200  {string}  string "Lesson created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/lessons/create [post]
func (h *Handler) CreateLesson(c *gin.Context) {
	var req pb.LessonCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	_, err := h.Clients.Lesson.CreateLesson(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create lesson:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Lesson created successfully")
}

// GetLesson retrieves a lesson by ID
// @Summary 		Get Lesson
// @Description 	Retrieve a lesson by its ID
// @Tags 			Lesson
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Lesson ID"
// @Success 		200  {object}  pb.LessonGet "Lesson details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Lesson not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/lessons/{id} [get]
func (h *Handler) GetLesson(c *gin.Context) {
	lessonID := c.Param("id")

	req := &pb.ById{Id: lessonID}

	res, err := h.Clients.Lesson.GetLesson(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get lesson:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateLesson updates an existing lesson
// @Summary 		Update Lesson
// @Description 	Update an existing lesson
// @Tags 			Lesson
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Lesson ID"
// @Param     		lesson body    pb.LessonUpt  true   "Lesson data"
// @Success 		200  {string}  string "Lesson updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/lessons/{id} [put]
func (h *Handler) UpdateLesson(c *gin.Context) {
	lessonID := c.Param("id")

	var body pb.LessonUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}
	req := pb.LessonUpdate{
		Id:   lessonID,
		Body: &body,
	}

	_, err := h.Clients.Lesson.UpdateLesson(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update lesson:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Lesson updated successfully")
}

// DeleteLesson deletes a lesson by ID
// @Summary 		Delete Lesson
// @Description 	Delete a lesson by its ID
// @Tags 			Lesson
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Lesson ID"
// @Success 		200  {string}  string "Lesson deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Lesson not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/lessons/{id} [delete]
func (h *Handler) DeleteLesson(c *gin.Context) {
	lessonID := c.Param("id")

	req := &pb.ById{Id: lessonID}

	_, err := h.Clients.Lesson.DeleteLesson(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete lesson:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Lesson deleted successfully")
}

// ListLessons retrieves a list of lessons
// @Summary List Lessons
// @Description Retrieve a list of lessons
// @Tags Lesson
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param course_id query string false "Course ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.LessonList "List of lessons"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /lessons/list [get]
func (h *Handler) ListLessons(c *gin.Context) {
	// Initialize the request and set default values
	var req pb.LessonFilter
	req.Filter = &pb.Pagination{}

	req.CourseId = c.Query("course_id")

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

	// Retrieve lessons from gRPC service
	res, err := h.Clients.Lesson.ListLessons(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list lessons:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

package handlers

import (
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateUserLesson creates a new user lesson
// @Summary       Create User Lesson
// @Description   Create a new user lesson
// @Tags          UserLesson
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         user_lesson body pb.UserLessonCreateReq true "User Lesson data"
// @Success       200  {string}  string "User Lesson created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /user-lessons/create [post]
func (h *Handler) CreateUserLesson(c *gin.Context) {
	var req pb.UserLessonCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.UserLesson.CreateUserLesson(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create user lesson:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User Lesson created successfully")
}

// GetUserLesson retrieves a user lesson by ID
// @Summary       Get User Lesson
// @Description   Retrieve a user lesson by its ID
// @Tags          UserLesson
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "User Lesson ID"
// @Success       200  {object}  pb.UserLesson "User Lesson details"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "User Lesson not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /user-lessons/{id} [get]
func (h *Handler) GetUserLesson(c *gin.Context) {
	userLessonID := c.Param("id")

	req := &pb.ById{Id: userLessonID}

	res, err := h.Clients.UserLesson.GetUserLesson(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get user lesson:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateUserLesson updates a user lesson by ID
// @Summary       Update User Lesson
// @Description   Update a user lesson
// @Tags          UserLesson
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "User Lesson ID"
// @Param         user_lesson body pb.UserLessonUpdateReq true "User Lesson update data"
// @Success       200  {string}  string "User Lesson updated successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /user-lessons/update/{id} [put]
func (h *Handler) UpdateUserLesson(c *gin.Context) {
	var req pb.UserLessonUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	req.Id = c.Param("id")

	_, err := h.Clients.UserLesson.UpdateUserLesson(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update user lesson:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User Lesson updated successfully")
}

// DeleteUserLesson deletes a user lesson by ID
// @Summary       Delete User Lesson
// @Description   Delete a user lesson by its ID
// @Tags          UserLesson
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "User Lesson ID"
// @Success       200  {string}  string "User Lesson deleted successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /user-lessons/delete/{id} [delete]
func (h *Handler) DeleteUserLesson(c *gin.Context) {
	userLessonID := c.Param("id")

	req := &pb.ById{Id: userLessonID}

	_, err := h.Clients.UserLesson.DeleteUserLesson(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete user lesson:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User Lesson deleted successfully")
}

// ListUserLessons retrieves a list of user lessons
// @Summary       List User Lessons
// @Description   Retrieve a list of user lessons with optional filters
// @Tags          UserLesson
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         user_course_id query string false "User Course ID"
// @Param         limit query int false "Limit"
// @Param         offset query int false "Offset"
// @Success       200  {object}  pb.UserLessonListsRes "List of user lessons"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /user-lessons/list [get]
func (h *Handler) ListUserLessons(c *gin.Context) {
	var req pb.UserLessonListsReq

	req.UserCourseId = c.Query("user_course_id")

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

	res, err := h.Clients.UserLesson.ListUserLessons(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list user lessons:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

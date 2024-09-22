package handlers

import (
	pb "muallimah-gateway/internal/pkg/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EnrollUserInCourse enrolls a user in a course
// @Summary        Enroll User in Course
// @Description    Enroll a user in a specified course
// @Tags           UserCourse
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          body body    pb.UserCourseCreateReq  true   "User and course IDs"
// @Success        200  {string}  string "User enrolled in course successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /usercourses/create [post]
func (h *Handler) EnrollUserInCourse(c *gin.Context) {
	var req pb.UserCourseCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.UserCourse.EnrollUserInCourse(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to enroll user in course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User enrolled in course successfully")
}

// GetUserCourse retrieves a user course by ID
// @Summary        Get User Course
// @Description    Retrieve a user course by its ID
// @Tags           UserCourse
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string  true  "User course ID"
// @Success        200  {object}  pb.UserCourse "User course details"
// @Failure        400  {string}  string "Invalid request"
// @Failure        404  {string}  string "User course not found"
// @Failure        500  {string}  string "Internal server error"
// @Router         /usercourses/{id} [get]
func (h *Handler) GetUserCourse(c *gin.Context) {
	courseID := c.Param("id")

	req := &pb.ById{Id: courseID}

	res, err := h.Clients.UserCourse.GetUserCourse(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get user course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateUserCourse updates a user course
// @Summary        Update User Course
// @Description    Update the details of an existing user course
// @Tags           UserCourse
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string      true   "User course ID"
// @Param          body body    pb.UserUpt  true   "User course update data"
// @Success        200  {string}  string "User course updated successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /usercourses/update/{id} [put]
func (h *Handler) UpdateUserCourse(c *gin.Context) {
	// Extract course ID from URL parameters
	courseID := c.Param("id")

	// Define a variable to hold the request body
	var body pb.UserUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		// Log the error and return a 400 Bad Request response if binding fails
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	// Create the request object for updating the user course
	req := pb.UserCourseUpdateReq{
		Id:   courseID,
		Body: &body,
	}

	// Call the gRPC method to update the user course
	_, err := h.Clients.UserCourse.UpdateUserCourse(c, &req)
	if err != nil {
		// Log the error and return a 500 Internal Server Error response if the gRPC call fails
		h.Logger.ERROR.Println("Failed to update user course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	// Return a 200 OK response if the update is successful
	c.JSON(200, "User course updated successfully")
}

// DeleteUserCourse deletes a user course
// @Summary        Delete User Course
// @Description    Delete a user course by its ID
// @Tags           UserCourse
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string  true  "User course ID"
// @Success        200  {string}  string "User course deleted successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /usercourses/delete/{id} [delete]
func (h *Handler) DeleteUserCourse(c *gin.Context) {
	courseID := c.Param("id")

	req := &pb.ById{Id: courseID}

	_, err := h.Clients.UserCourse.DeleteUserCourse(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete user course:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User course deleted successfully")
}

// ListUserCourses lists all user courses with optional filters
// @Summary        List User Courses
// @Description    List all user courses with optional filters and pagination
// @Tags           UserCourse
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          user_id query   string  false  "User ID"
// @Param          limit query string false "Limit"
// @Param          offset query string false "Offset"
// @Success        200  {object}  pb.UserCourseListsRes "List of user courses"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"“
// @Router         /usercourses/list [get]
func (h *Handler) ListUserCourses(c *gin.Context) {
	var req pb.UserCourseListsReq
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

	res, err := h.Clients.UserCourse.ListUserCourses(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list user courses:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

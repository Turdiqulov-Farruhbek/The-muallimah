package handlers

import (
	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user
// @Summary        Create User
// @Description    Create a new user
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          user body    pb.UserCreateReqForSwagger  true   "User data"
// @Success        200  {string}  string "User created successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/create [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req pb.UserCreateReqForSwagger
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.User.CreateUser(c, &pb.UserCreateReq{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Dob:         req.Dob,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Occupation:  req.Occupation,
		Address:     req.Address,
		Password:    req.Password,
		Gender:      req.Gender,
	})
	if err != nil {
		h.Logger.ERROR.Println("Failed to create user:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User created successfully")
}

// GetUserByID retrieves a user by ID
// @Summary        Get User by ID
// @Description    Retrieve a user by their ID
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string  true  "User ID"
// @Success        200  {object}  pb.UserGetRes "User details"
// @Failure        400  {string}  string "Invalid request"
// @Failure        404  {string}  string "User not found"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	req := &pb.ById{Id: userID}

	res, err := h.Clients.User.GetUserByID(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get user:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// GetUserByEmail retrieves a user by email
// @Summary        Get User by Email
// @Description    Retrieve a user by their email
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          email query   string  true  "User email"
// @Success        200  {object}  pb.UserGetRes "User details"
// @Failure        400  {string}  string "Invalid request"
// @Failure        404  {string}  string "User not found"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/by-email [get]
func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	req := &pb.ByEmail{Email: email}

	res, err := h.Clients.User.GetUserByEmail(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get user by email:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// UpdateUser updates an existing user
// @Summary        Update User
// @Description    Update an existing user
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string       true   "User ID"
// @Param          user body    pb.UserUpdateReq  true   "User data"
// @Success        200  {string}  string "User updated successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var req pb.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.User.UpdateUser(c, &pb.UserUpdateReq{
		Id:   userID,
		Body: &req,
	})
	if err != nil {
		h.Logger.ERROR.Println("Failed to update user:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User updated successfully")
}

// ChangeUserPassword changes a user's password
// @Summary        Change User Password
// @Description    Change the password for a user
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          body body    pb.UserRecoverPasswordReq  true   "User email and new password"
// @Success        200  {string}  string "Password changed successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/change-password [post]
func (h *Handler) ChangeUserPassword(c *gin.Context) {
	var req pb.UserRecoverPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.User.ChangeUserPassword(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to change password:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Password changed successfully")
}

// ChangeUserPFP updates a user's profile picture
// @Summary        Change User Profile Picture
// @Description    Update the profile picture of a user
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          body body    pb.UserCreatePFPReqForSwagger  true   "User photo URL"
// @Success        200  {string}  string "Profile picture updated successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/change-photo [post]
func (h *Handler) ChangeUserPFP(c *gin.Context) {
	var req pb.UserCreatePFPReqForSwagger
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.User.ChangeUserPFP(c, &pb.UserChangePFPReq{
		PhotoUrl: req.PhotoUrl,
	})
	if err != nil {
		h.Logger.ERROR.Println("Failed to update profile picture:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Profile picture updated successfully")
}

// IsEmailExists checks if an email exists
// @Summary        Check Email Exists
// @Description    Check if the provided email already exists
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          email query   string  true  "Email address"
// @Success        200  {object}  pb.UserEmailCheckRes "Email existence status"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/email-exists [get]
func (h *Handler) IsEmailExists(c *gin.Context) {
	email := c.Query("email")

	req := &pb.UserEmailCheckReq{Email: email}

	res, err := h.Clients.User.IsEmailExists(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to check email existence:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// ConfirmUser confirms a user's email
// @Summary        Confirm User Email
// @Description    Confirm the email address of a user
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          email query   string  true  "User email"
// @Param          code query    string  true  "Confirmation code"
// @Success        200  {string}  string "User email confirmed successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/confirm [post]
func (h *Handler) ConfirmUser(c *gin.Context) {
    email := c.Query("email")

    // Create the request with the correct type
    req := &pb.ByEmail{Email: email}

    // Call the gRPC method with the correct request type
    _, err := h.Clients.User.ConfirmUser(c, req)
    if err != nil {
        h.Logger.ERROR.Println("Failed to confirm user:", err)
        c.JSON(500, "Internal server error: "+err.Error())
        return
    }

    c.JSON(200, "User email confirmed successfully")
}


// DeleteUser deletes a user
// @Summary        Delete User
// @Description    Delete a user by their ID
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          id path    string  true  "User ID"
// @Success        200  {string}  string "User deleted successfully"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	req := &pb.ById{Id: userID}

	_, err := h.Clients.User.DeleteUser(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete user:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "User deleted successfully")
}

// ListUsers lists all users with optional filters
// @Summary        List Users
// @Description    List all users with optional filters and pagination
// @Tags           User
// @Accept         json
// @Produce        json
// @Security       BearerAuth
// @Param          filters query   pb.UserGetAllFilter  true  "Filters for user listing"
// @Param          pagination query   pb.Pagination  true  "Pagination data"
// @Success        200  {object}  pb.UsersGetAllRes "List of users"
// @Failure        400  {string}  string "Invalid request"
// @Failure        500  {string}  string "Internal server error"
// @Router         /users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	var req pb.UsersGetAllReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind query parameters:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	res, err := h.Clients.User.ListUsers(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list users:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

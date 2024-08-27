package handlers

import (
	"auth-service/internal/pkg/config"
	pb "auth-service/internal/pkg/genproto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// ForgotPassword godoc
// @Summary Forgot passwrod
// @Description Sends a confirmation code to email recovery password
// @Tags password-recovery
// @Accept json
// @Produce json
// @Param credentials body pb.ByEmail true "User login credentials"
// @Success 200 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Page not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /forgot-password [POST]
func (h *HTTPHandler) ForgotPassword(c *gin.Context) {
	var req pb.ByEmail
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	user, err := h.US.GetUserByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = config.SendConfirmationCode(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code to email", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code sent to your email. Please use your code within 3 minutes."})
}

// RecoverPassword godoc
// @Summary Recover password (Use this one after sending verification code)
// @Description Verifies the code and updates the password
// @Tags password-recovery
// @Accept json
// @Produce json
// @Param request body pb.UserChangePasswordReq true "Recover Password Request"
// @Success 200 {object} string "Password successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Failure 500 {object} string "Error updating password"
// @Router /recover-password [post]
func (h *HTTPHandler) RecoverPassword(c *gin.Context) {
	var req pb.UserChangePasswordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, code, and new password are required fields."})
		return
	}

	if err := config.IsValidPassword(req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	storedCode, err := h.RDB.Get(context.Background(), req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code expired or email not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email not found in a recovery requests!"})
		return
	}

	if storedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect verification code"})
		return
	}

	hashedPassword, err := config.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't has your password", "details": err.Error()})
		return
	}
	_, err = h.US.ChangeUserPassword(c, &pb.UserChangePasswordReq{Email: req.Email, NewPassword: hashedPassword})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password", "details": err.Error()})
		return
	}
	h.RDB.Del(context.Background(), req.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully updated"})
}

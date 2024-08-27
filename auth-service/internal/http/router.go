package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/internal/http/docs"
	"auth-service/internal/http/handlers"
	"auth-service/internal/http/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(h *handlers.HTTPHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)
	router.POST("/set-pfp", h.SetPFP)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)

	router.GET("/user/:id", h.GetByID)

	return router
}

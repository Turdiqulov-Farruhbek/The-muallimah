package http

import (
	_ "muallimah-gateway/docs"
	"muallimah-gateway/internal/http/handlers"
	"muallimah-gateway/internal/http/middlerware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Finance Tracker API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	enforcer, err := casbin.NewEnforcer("./internal/http/casbin/model.conf", "./internal/http/casbin/policy.csv")
	if err != nil {
		panic(err)
	}
	router.Use(middlerware.NewAuth(enforcer))

	materials := router.Group("/materials")
	{
		materials.POST("/create", h.CreateMaterial)
		materials.GET("/:id", h.GetMaterial)
		materials.PUT("/update/:id", h.UpdateMaterial)
		materials.DELETE("/delete/:id", h.DeleteMaterial)
		materials.GET("/list", h.ListMaterials)
	}

	books := router.Group("/books")
	{
		books.POST("/create", h.CreateBook)
        books.GET("/:id", h.GetBook)
        books.PUT("/update/:id", h.UpdateBook)
        books.DELETE("/delete/:id", h.DeleteBook)
        books.GET("/list", h.ListBooks)
	}

	categories := router.Group("/categories")
	{
		categories.POST("/create", h.CreateCategory)
        categories.GET("/:id", h.GetCategory)
        categories.PUT("/update/:id", h.UpdateCategory)
        categories.DELETE("/delete/:id", h.DeleteCategory)
        categories.GET("/list", h.ListCategories)
	}

	certificates := router.Group("/certificates")
	{
		certificates.POST("/create", h.CreateCertificate)
        certificates.GET("/:id", h.GetCertificate)
        certificates.PUT("/update/:id", h.UpdateCertificate)
        certificates.DELETE("/delete/:id", h.DeleteCertificate)
        certificates.GET("/list", h.ListCertificates)
	}

	courses := router.Group("/course")
	{
		courses.POST("/create", h.CreateCourse)
        courses.GET("/:id", h.GetCourse)
        courses.PUT("/update/:id", h.UpdateCourse)
        courses.DELETE("/delete/:id", h.DeleteCourse)
        courses.GET("/list", h.ListCourses)
	}

	evalution := router.Group("/evalution")
	{
		evalution.POST("/create", h.CreateEvaluation)
        evalution.GET("/:id", h.GetEvaluation)
        evalution.DELETE("/delete/:id", h.DeleteEvaluation)
        evalution.GET("/list", h.ListEvaluations)
	}

	lessons := router.Group("/lessons")
	{
		lessons.POST("/create", h.CreateLesson)
        lessons.GET("/:id", h.GetLesson)
        lessons.PUT("/update/:id", h.UpdateLesson)
        lessons.DELETE("/delete/:id", h.DeleteLesson)
        lessons.GET("/list", h.ListLessons)
	}

	// notifications := router.Group("/notification")
	// {
	// 	notifications.POST("/create", h.CreateNotification)
    //     notifications.GET("/:id", h.GetNotification)
    //     notifications.PUT("/update/:id", h.UpdateNotification)
    //     notifications.DELETE("/delete/:id", h.DeleteNotification)
    //     notifications.GET("/list", h.ListNotifications)
	// }

	orders := router.Group("/orders")
	{
		orders.POST("/create", h.CreateOrder)
        orders.GET("/:id", h.GetOrder)
        orders.PUT("/update/:id", h.UpdateOrder)
        orders.DELETE("/delete/:id", h.DeleteOrder)
        orders.GET("/list", h.ListOrders)
	}

	posts := router.Group("/posts")
	{
		posts.POST("/create", h.CreatePost)
        posts.GET("/:id", h.GetPost)
        posts.PUT("/update/:id", h.UpdatePost)
        posts.DELETE("/delete/:id", h.DeletePost)
        posts.GET("/list", h.GetPosts)
		posts.POST("/pictures/add", h.AddPicturePost)
		posts.POST("/pictures/delete", h.DeletePicturePost)
	}

	products := router.Group("/products")
{
	products.POST("/create", h.CreateProduct)
	products.GET("/:id", h.GetProduct)
	products.PUT("/update/:id", h.UpdateProduct)
	products.DELETE("/delete/:id", h.DeleteProduct)
	products.GET("/list", h.ListProducts)
	products.POST("/pictures/add", h.AddPictureProduct)
	products.POST("/pictures/delete", h.DeletePictureProduct)
}


	transactions := router.Group("/transactions")
	{
		transactions.POST("/create", h.CreateTransaction)
        transactions.GET("/:id", h.GetTransaction)
        transactions.GET("/Get balance/:id", h.GetBalance)
        transactions.GET("/list", h.ListTransactions)
	}

	userlessons := router.Group("/userlessons")
	{
		userlessons.POST("/create", h.CreateUserLesson)
        userlessons.GET("/:id", h.GetUserLesson)
        userlessons.PUT("/update/:id", h.UpdateUserLesson)
        userlessons.DELETE("/delete/:id", h.DeleteUserLesson)
        userlessons.GET("/list", h.ListUserLessons)
	}

	usercourses := router.Group("/usercourses")
	{
		usercourses.POST("/create", h.EnrollUserInCourse)
        usercourses.GET("/:id", h.GetUserCourse)
        usercourses.PUT("/update/:id", h.UpdateUserCourse)
        usercourses.DELETE("/delete/:id", h.DeleteUserCourse)
        usercourses.GET("/list", h.ListUserCourses)
	}

	users := router.Group("/users")
	{
		users.POST("/create", h.CreateUser)
		users.GET("/:id", h.GetUserByID)
		users.GET("/by-email", h.GetUserByEmail)
		users.PUT("/:id", h.UpdateUser)
		users.POST("/change-password", h.ChangeUserPassword)
		users.POST("/change-photo", h.ChangeUserPFP)
		users.GET("/email-exists", h.IsEmailExists)
		users.POST("/confirm", h.ConfirmUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("/", h.ListUsers)
	}

	return router
}

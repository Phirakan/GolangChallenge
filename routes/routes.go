package routes

import (
	"net/http"
	"time"

	"golang-test/controllers"
	"golang-test/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	
	router.Use(middleware.LoggerMiddleware())

		router.GET("/api/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "golang-api",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.GET("/users", controllers.GetAllUsers)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
	}
}
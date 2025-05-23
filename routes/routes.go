package routes

import (
	"github.com/gin-gonic/gin"
	"golang-test/controllers"
	"golang-test/middleware"
)

func SetupRoutes(router *gin.Engine) {
	
	router.Use(middleware.LoggerMiddleware())

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
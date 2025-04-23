package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Get tasks - authenticated"})
		})

		auth.POST("/tasks", func(c *gin.Context) {
			var task models.Task

			if err := c.ShouldBindJSON(&task); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			username := c.MustGet("username").(string)
			task.User = username

			created := data.CreateTask(task)
			c.JSON(201, created)
		})

		admin := auth.Group("/admin")
		admin.Use(middleware.AdminOnly())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Admin dashboard"})
			})
		}
	}

	return r
}
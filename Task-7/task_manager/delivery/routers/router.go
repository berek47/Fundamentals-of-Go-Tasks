package routers

import (
	"task_manager/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	userRepo := repositories.NewInMemoryUserRepo()
	taskRepo := repositories.NewInMemoryTaskRepo()

	userUsecase := &usecases.UserUsecase{UserRepo: userRepo}
	taskUsecase := &usecases.TaskUsecase{TaskRepo: taskRepo}

	userController := &controllers.UserController{Usecase: userUsecase}
	taskController := &controllers.TaskController{Usecase: taskUsecase}

	r := gin.Default()
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware())
	{
		auth.GET("/tasks", taskController.GetTasks)
		auth.POST("/tasks", taskController.CreateTask)

		admin := auth.Group("/admin")
		admin.Use(infrastructure.AdminOnly())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Admin dashboard"})
			})
		}
	}

	return r
}
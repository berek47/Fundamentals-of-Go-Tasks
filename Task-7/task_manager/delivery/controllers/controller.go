package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase *usecases.UserUsecase
}

type TaskController struct {
	Usecase *usecases.TaskUsecase
}

func (u *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if req.Role == "" {
		req.Role = "user"
	}
	user, err := u.Usecase.Register(req.Username, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (u *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, err := u.Usecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, _ := infrastructure.GenerateToken(user.Username, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	username := c.MustGet("username").(string)
	task.User = username
	created := t.Usecase.Create(task)
	c.JSON(http.StatusCreated, created)
}

func (t *TaskController) GetTasks(c *gin.Context) {
	tasks := t.Usecase.GetAll()
	c.JSON(http.StatusOK, tasks)
}
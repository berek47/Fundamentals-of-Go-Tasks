package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.InitMongoDB("mongodb+srv://berek47:mysecurepassword@cluster0.bxssrvt.mongodb.net/taskdb?retryWrites=true&w=majority&appName=Cluster0", "taskdb", "tasks")

    r := router.SetupRouter()
    r.Run(":8080")
}
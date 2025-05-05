package main

import (
	"task_manager/delivery/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
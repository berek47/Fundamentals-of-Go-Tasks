package main

import (
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

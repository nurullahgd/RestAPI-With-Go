package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nurullahgd/RestAPI-With-Go/config"
	"github.com/nurullahgd/RestAPI-With-Go/routes"
)

func main() {
	fmt.Println("Test")
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":4569")
}

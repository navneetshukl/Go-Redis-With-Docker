package main

import (
	"go_modules/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", routes.Home)
	router.POST("/push", routes.Push)
	router.Run()

}

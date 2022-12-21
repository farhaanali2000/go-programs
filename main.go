package main

import (
	"github.com/gin-gonic/gin"
	"www.github.com/farhaan/go-programs/controllers"
)

func main() {
	// fmt.Printf("Goal")

	r := gin.Default()

	r.GET("/home/:code", controllers.HomeHandler)

	r.POST("/submit", controllers.AddMovie)

	r.Run(":4040")
}

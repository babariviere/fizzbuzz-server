package main

import (
	"fmt"

	_ "github.com/babariviere/fizzbuzz-server/docs"
	"github.com/babariviere/fizzbuzz-server/pkg/fizzbuzz"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Fizzbuzz API server
// @version     1.0
// @description API to generate fizzbuzz sequences
// @BasePath    /
func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	{
		handler := fizzbuzz.NewHandler()
		fizzbuzz.RegisterRoutes(&r.RouterGroup, &handler)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":3000")
}

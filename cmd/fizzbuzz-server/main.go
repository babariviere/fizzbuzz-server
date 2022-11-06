package main

import (
	"fmt"

	"github.com/babariviere/fizzbuzz-server/pkg/fizzbuzz"
    _ "github.com/babariviere/fizzbuzz-server/docs"
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

    fizzbuzzH := fizzbuzz.Handler{Service: fizzbuzz.NewFizzbuzz()}
    fizzbuzz.RegisterRoutes(&r.RouterGroup, &fizzbuzzH)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    r.Run(":3000")
}

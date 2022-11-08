package main

import (
	"fmt"
	"log"

	_ "github.com/babariviere/fizzbuzz-server/docs"
	"github.com/babariviere/fizzbuzz-server/pkg/fizzbuzz"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
    HttpAddress string `mapstructure:"HTTP_ADDRESS"`
    HttpPort int `mapstructure:"HTTP_PORT"`
}

// @title       Fizzbuzz API server
// @version     1.0
// @description API to generate fizzbuzz sequences
// @BasePath    /
func main() {
    cfg, err := LoadConfig(".")
    if err != nil {
        log.Fatalf("cannot read config: %v", err)
    }

	r := gin.Default()
	{
		handler := fizzbuzz.NewHandler()
		fizzbuzz.RegisterRoutes(&r.RouterGroup, &handler)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(fmt.Sprintf("%v:%v", cfg.HttpAddress, cfg.HttpPort))
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetDefault("HTTP_ADDRESS", "127.0.0.1")
    viper.SetDefault("HTTP_PORT", "3000")
    viper.SetConfigName("")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return
	}

    err = viper.Unmarshal(&config)
    return
}
